package render

import (
	"io"
)

type (
	IRenderer interface {
		Render(io.Writer, string, any) error
	}
)
