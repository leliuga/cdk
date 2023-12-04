package database

import (
	"bytes"
	"database/sql/driver"
	"fmt"
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
func (d *Dialect) String() string {
	return DialectNames[*d]
}

// Bytes returns the Dialect as a []byte.
func (d *Dialect) Bytes() []byte {
	return []byte(d.String())
}

// Value outputs the Dialect as a value.
func (d *Dialect) Value() (driver.Value, error) {
	return d.String(), nil
}

// MarshalJSON outputs the Dialect as a json.
func (d *Dialect) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

// UnmarshalJSON parses the Dialect from json.
func (d *Dialect) UnmarshalJSON(data []byte) error {
	v, err := ParseDialect(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*d = *v

	return nil
}

// Validate returns true if the Dialect is valid.
func (d *Dialect) Validate() bool {
	return *d != DialectInvalid
}

// ParseDialect parses the Dialect from string.
func ParseDialect(value string) (*Dialect, error) {
	value = strings.ToLower(value)
	for k, v := range DialectNames {
		if v == value {
			return &k, nil
		}
	}

	return nil, fmt.Errorf("unsupported dialect: %s", value)
}
