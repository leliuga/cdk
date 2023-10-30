package http

type (
	// Header defines an HTTP header.
	Header uint16

	// Method defines an HTTP method.
	Method uint8

	// Status defines an HTTP status.
	Status uint16

	// Headers represents a collection of HTTP headers.
	Headers map[Header]string

	// Error represents an HTTP error.
	Error struct {
		Status  Status `json:"status" yaml:"Status"`
		Message string `json:"message" yaml:"Message"`
	}

	// IError defines the interface for an HTTP error.
	IError interface {
		error

		// StatusCode returns the HTTP status code of the error
		StatusCode() Status
	}
)
