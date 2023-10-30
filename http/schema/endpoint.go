package schema

import (
	"fmt"
	"strings"

	"github.com/leliuga/cdk/constants"
	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/types"
	"github.com/leliuga/cdk/validation"
)

func NewEndpoint(name string, method http.Method, path string) *Endpoint {
	return &Endpoint{
		Name:    name,
		Method:  method,
		Path:    path,
		Headers: http.Headers{},
		Expect:  NewExpect(),
	}
}

// Validate makes Endpoint validatable by implementing [validation.Validatable] interface.
func (e *Endpoint) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Name, validation.Required, validation.Length(1, 63), validation.Match(constants.NameRegex).Error(constants.InvalidName)),
		validation.Field(&e.Method, validation.Required, validation.In(validation.ToAnySliceFromMapKeys(http.MethodNames)...).Error(fmt.Sprintf("A method value must be one of: %s", strings.Join(types.ToMap(http.MethodNames).Values(), ", ")))),
		validation.Field(&e.Path, validation.Required),
	)
}
