package database

import (
	"bytes"
	"strings"
)

const (
	DialectInvalid Dialect = iota + 1 //
	DialectGraphQL
	DialectMysql
	DialectPostgreSQL
	DialectClickhouse
	DialectBigQuery
)

var (
	DialectNames = map[Dialect]string{
		DialectGraphQL:    "graphql",
		DialectMysql:      "mysql",
		DialectPostgreSQL: "postgresql",
		DialectClickhouse: "clickhouse",
		DialectBigQuery:   "bigquery",
	}
)

// String outputs the Dialect as a string.
func (d Dialect) String() string {
	return DialectNames[d]
}

// MarshalJSON outputs the Dialect as a json.
func (d Dialect) MarshalJSON() ([]byte, error) {
	if !d.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + d.String() + `"`), nil
}

// UnmarshalJSON parses the Dialect from json.
func (d *Dialect) UnmarshalJSON(data []byte) error {
	str := string(bytes.Trim(data, `"`))
	if dialect := ParseDialect(str); dialect.Validate() {
		*d = dialect
	}

	return nil
}

// Validate returns true if the Dialect is valid.
func (d Dialect) Validate() bool {
	return d != DialectInvalid
}

// ParseDialect parses the Dialect from string.
func ParseDialect(value string) Dialect {
	value = strings.ToLower(value)
	for k, v := range DialectNames {
		if v == value {
			return k
		}
	}

	return DialectInvalid
}
