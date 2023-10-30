package schema

import (
	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/types"
	"github.com/leliuga/validation"
)

type (
	// Endpoint is an HTTP client endpoint.
	Endpoint struct {
		validation.Validatable `json:"-"`
		Name                   string            `json:"name"`
		Method                 http.Method       `json:"method"`
		Path                   string            `json:"path"`
		Description            string            `json:"description"`
		Documentation          string            `json:"documentation"`
		Deprecated             string            `json:"deprecated"`
		Labels                 types.Map[string] `json:"labels"`
		Headers                http.Headers      `json:"headers"`
		Payload                any               `json:"payload"`
		Expect                 *Expect           `json:"expect"`
	}

	// Expect is an HTTP response expectation.
	Expect struct {
		Status  http.Status  `json:"status"`
		Headers http.Headers `json:"headers"`
	}

	// Endpoints represents the http client endpoints.
	Endpoints []*Endpoint
)
