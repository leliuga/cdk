package database

import (
	"github.com/leliuga/cdk/types"
)

// NewColumn creates a new Column instance
func NewColumn(t types.Type, name, description, nativeType string) *Column {
	return &Column{
		Name:        name,
		Description: description,
		Type:        &t,
		NativeType:  nativeType,
		Creatable:   true,
		Updatable:   true,
		Readable:    true,
	}
}
