// Package gql provides a GraphQL schema to Go struct code generator.
package gql

import (
	"github.com/leliuga/cdk/types"
	"github.com/vektah/gqlparser/v2/ast"
)

type (
	Item struct {
		Implements  []byte
		Definitions []byte
	}

	Schema2Go struct {
		Schema *ast.Schema
		Items  types.Map[Item]
	}
)
