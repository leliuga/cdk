package configurator

import (
	"encoding/base64"

	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
	"github.com/google/brotli/go/cbrotli"
)

// NewEnvironment creates a new Environment instance.
func NewEnvironment(name, description string) *Environment {
	return &Environment{
		Name:        name,
		Description: description,
		Entries:     []*EnvironmentEntry{},
	}
}

// Set adds an EnvironmentEntry to the Environment.
func (e *Environment) Set(entry *EnvironmentEntry) {
	e.Entries = append(e.Entries, entry)
}

// Get returns an EnvironmentEntry by key.
func (e *Environment) Get(key string) (*EnvironmentEntry, bool) {
	for _, entry := range e.Entries {
		if entry.Key == key {
			return entry, true
		}
	}

	return &EnvironmentEntry{}, false
}

// Delete removes an EnvironmentEntry by key.
func (e *Environment) Delete(key string) {
	for i, entry := range e.Entries {
		if entry.Key == key {
			e.Entries = append(e.Entries[:i], e.Entries[i+1:]...)
			return
		}
	}
}

// Marshal returns the Environment as a string in the given format.
func (e *Environment) Marshal(format string) string {
	switch format {
	case "json":
		if out, err := json.Marshal(&e); err == nil {
			return string(out) + "\n"
		}
	case "yaml":
		if out, err := yaml.MarshalWithOptions(&e, yaml.UseJSONMarshaler()); err == nil {
			return string(out)
		}
	case "dotenv":
		var out string
		for _, entry := range e.Entries {
			out += entry.Key + "=" + entry.Value + "\n"
		}
		return out
	case "init":
		b := []byte(e.Marshal("dotenv"))
		if out, err := cbrotli.Encode(b, cbrotli.WriterOptions{Quality: 9}); err == nil {
			return base64.StdEncoding.EncodeToString(out) + "\n"
		}
	}

	return ""
}
