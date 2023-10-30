package database

import (
	"github.com/leliuga/cdk/types"
)

// NewTable creates a new Table instance
func NewTable(name, description string, columns ...*Column) *Table {
	table := &Table{
		Name:        name,
		Description: description,
		Columns:     types.NewMap[*Column](),
	}

	for _, column := range columns {
		table.Columns[column.Name] = column
	}

	return table
}
