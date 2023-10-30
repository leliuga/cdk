package schema

import (
	"github.com/leliuga/cdk/http"
	"github.com/pkg/errors"
)

// NewExpect creates a new Expect.
func NewExpect() *Expect {
	return &Expect{
		Status:  http.StatusOK,
		Headers: http.Headers{},
	}
}

// Validate validates the given response.
func (e *Expect) Validate(status http.Status, headers http.Headers) error {
	if err := e.ValidateStatus(status); err != nil {
		return err
	}

	if err := e.ValidateHeader(headers); err != nil {
		return err
	}

	return nil
}

// ValidateStatus validates the status code of the given response.
func (e *Expect) ValidateStatus(status http.Status) error {
	if status != e.Status {
		return errors.Errorf("unexpected status code %d, expected %d.", status, e.Status)
	}

	return nil
}

// ValidateHeader validates the headers of the given response.
func (e *Expect) ValidateHeader(headers http.Headers) error {
	for k, v := range e.Headers {
		if value := headers[k]; value != v {
			return errors.Errorf("unexpected header %s=%s, expected %s.", k, value, v)
		}
	}

	return nil
}
