package html

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/gofiber/storage/memory"
	"golang.org/x/exp/maps"
)

const (
	Extension      = ".htm"
	BlocksPrefix   = "blocks/"
	PartialsPrefix = "partials/"
	LayoutsPrefix  = "layouts/"
	PagesPrefix    = "pages/"

	DefaultLanguage     = "en"
	DefaultLayout       = "default"
	DefaultRobotsIndex  = true
	DefaultRobotsFollow = true
)

// NewRender creates a new Html from the given directory.
func NewRender(options *Options) *Html {
	h := &Html{
		Options:  options,
		Blocks:   Blocks{},
		Partials: Partials{},
		Layouts:  Layouts{},
		Pages:    Pages{},
		mutex:    &sync.Mutex{},
	}

	h.Load()

	h.storage = memory.New(memory.Config{
		GCInterval: 1 * time.Minute,
	})

	return h
}

// Load loads all templates from the given directory.
func (h *Html) Load() error {
	set := pongo2.NewSet(Extension, pongo2.MustNewLocalFileSystemLoader(h.Directory))
	pongo2.SetAutoescape(false)

	for name, filter := range h.filters {
		if err := pongo2.RegisterFilter(name, filter); err != nil {
			return err
		}
	}

	return filepath.Walk(h.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info == nil || info.IsDir() || len(Extension) >= len(path) || path[len(path)-len(Extension):] != Extension {
			return nil
		}

		rel, err := filepath.Rel(h.Directory, path)
		if err != nil {
			return err
		}

		template := strings.TrimSuffix(filepath.ToSlash(rel), Extension)
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		minify := h.Minify
		switch {
		case strings.HasPrefix(template, BlocksPrefix):
			if h.Blocks[template], err = NewBlock(set, strings.TrimPrefix(template, BlocksPrefix), content, minify); err != nil {
				return err
			}
		case strings.HasPrefix(template, PartialsPrefix):
			if h.Partials[template], err = NewPartial(set, strings.TrimPrefix(template, PartialsPrefix), content, minify); err != nil {
				return err
			}
		case strings.HasPrefix(template, LayoutsPrefix):
			if h.Layouts[template], err = NewLayout(set, strings.TrimPrefix(template, LayoutsPrefix), content, minify); err != nil {
				return err
			}
		case strings.HasPrefix(template, PagesPrefix):
			if h.Pages[template], err = NewPage(set, strings.TrimPrefix(template, PagesPrefix), content, minify); err != nil {
				return err
			}
		}

		h.loaded = true

		return nil
	})
}

// Render renders the given template with the given variables.
func (h *Html) Render(writer io.Writer, template string, variables any, l ...string) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	var vars = h.Variables

	if v, ok := variables.(pongo2.Context); ok {
		maps.Copy(vars, v)
	}

	cacheKey := h.cacheKey(template, vars)
	if h.Debug {
		h.loaded = false
		if err := h.Load(); err != nil {
			return err
		}
	}

	if h.Cache && h.storage != nil {
		if cached, err := h.storage.Get(cacheKey); len(cached) > 0 && err == nil {
			_, err = writer.Write(cached)
			return nil
		}
	}

	var buffer bytes.Buffer
	switch {
	case strings.HasPrefix(template, BlocksPrefix):
		block, found := h.Blocks[template]
		if !found {
			return fmt.Errorf("block with name '%s' could not be found", strings.TrimPrefix(template, BlocksPrefix))
		}

		if err := block.Render(&buffer, vars); err != nil {
			return err
		}
	case strings.HasPrefix(template, PartialsPrefix):
		partial, found := h.Partials[template]
		if !found {
			return fmt.Errorf("partial with name '%s' could not be found", strings.TrimPrefix(template, PartialsPrefix))
		}

		if err := partial.Render(&buffer, vars); err != nil {
			return err
		}
	case strings.HasPrefix(template, LayoutsPrefix):
		layout, found := h.Layouts[template]
		if !found {
			return fmt.Errorf("layout with name '%s' could not be found", strings.TrimPrefix(template, LayoutsPrefix))
		}

		if err := layout.Render(&buffer, vars); err != nil {
			return err
		}
	case strings.HasPrefix(template, PagesPrefix):
		page, found := h.Pages[template]
		if !found {
			return fmt.Errorf("page with name '%s' could not be found", strings.TrimPrefix(template, PagesPrefix))
		}

		layout, found := h.Layouts[LayoutsPrefix+page.Layout]
		if !found {
			return fmt.Errorf("in the page '%s' layout with name '%s' could not be found", strings.TrimPrefix(template, PagesPrefix), strings.TrimPrefix(page.Layout, LayoutsPrefix))
		}

		var buf strings.Builder
		if err := page.Render(&buf, vars); err != nil {
			return err
		}

		vars["embed"] = buf.String()

		if err := layout.Render(&buffer, vars); err != nil {
			return err
		}
	}

	rendered := buffer.Bytes()
	if h.Cache && h.storage != nil {
		if err := h.storage.Set(cacheKey, rendered, 1*time.Minute); err != nil {
			return err
		}
	}

	_, err := writer.Write(rendered)

	return err
}

// cacheKey returns a unique key for the given template and variables.
func (h *Html) cacheKey(template string, variables pongo2.Context) string {
	data := append([]byte(template), []byte(fmt.Sprintf("%v", variables))...)
	return fmt.Sprintf("html-%x", md5.Sum(data))
}
