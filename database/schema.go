package database

import (
	"github.com/leliuga/cdk/types"
)

// NewSchema creates a new Schema instance
func NewSchema(name, description string, tables ...*Table) *Schema {
	schema := &Schema{
		Name:        name,
		Description: description,
		Tables:      types.NewMap[*Table](),
	}

	for _, table := range tables {
		schema.Tables[table.Name] = table
	}

	return schema
}
