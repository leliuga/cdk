// Package database provides a database abstraction layer for SQL and NoSQL databases.
package database

import (
	"time"

	"github.com/leliuga/cdk/types"
)

type (
	// Options represents the storage options.
	Options struct {
		SourcesDsn            types.Map[*types.URI] `json:"sources_dsn"              env:"SOURCES_DSN"`
		ReplicasDsn           types.Map[*types.URI] `json:"replicas_dsn"             env:"REPLICAS_DSN"`
		MaxOpenConnections    int                   `json:"max_open_connections"     env:"MAX_OPEN_CONNECTIONS"`
		MaxIdleConnections    int                   `json:"max_idle_connections"     env:"MAX_IDLE_CONNECTIONS"`
		MaxLifetimeConnection time.Duration         `json:"max_lifetime_connection"  env:"MAX_LIFETIME_CONNECTION"`
		MaxIdleTimeConnection time.Duration         `json:"max_idle_time_connection" env:"MAX_IDLE_TIME_CONNECTION"`
		Options               types.Map[string]     `json:"options"`
	}

	// Schema defines a single storage schema structure.
	Schema struct {
		Name          string            `json:"name"`
		Description   string            `json:"description"`
		Deprecated    string            `json:"deprecated"`
		Replacement   string            `json:"replacement"`
		Documentation string            `json:"documentation"`
		Collation     string            `json:"collation"`
		Tables        types.Map[*Table] `json:"tables"`
	}

	// Table defines a single Schema table structure.
	Table struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Deprecated  string             `json:"deprecated"`
		Replacement string             `json:"replacement"`
		Type        string             `json:"type"`
		Engine      string             `json:"engine"`
		RowFormat   string             `json:"row_format"`
		Codec       string             `json:"codec"`
		Collation   string             `json:"collation"`
		ReadOnly    bool               `json:"read_only"`
		Columns     types.Map[*Column] `json:"columns"`
	}

	// Column defines a single Table field structure.
	Column struct {
		Name          string      `json:"name"`
		Description   string      `json:"description"`
		Deprecated    string      `json:"deprecated"`
		Replacement   string      `json:"replacement"`
		Type          *types.Type `json:"type"`
		NativeType    string      `json:"native_type"`
		Codec         string      `json:"codec"`
		Collation     string      `json:"collation"`
		Default       string      `json:"default"`
		Validation    string      `json:"validation"`
		Sensitive     bool        `json:"sensitive"`
		AutoIncrement bool        `json:"auto_increment"`
		Primary       bool        `json:"primary"`
		Index         bool        `json:"index"`
		Unique        bool        `json:"unique"`
		Nullable      bool        `json:"nullable"`
		Creatable     bool        `json:"creatable"`
		Updatable     bool        `json:"updatable"`
		Readable      bool        `json:"readable"`
	}

	// Version represents a storage version
	Version struct {
		Database       string   `json:"database"`
		Dialect        *Dialect `json:"dialect"`
		Version        string   `json:"version"`
		Description    string   `json:"description"`
		CompileMachine string   `json:"compile_machine"`
		CompileOs      string   `json:"compile_os"`
	}

	// Dialect represents the storage dialect.
	Dialect uint8

	// Option represents the storage option.
	Option func(*Options)

	// IKeyValue interface for communicating with different database/key-value providers
	IKeyValue interface {
		// Get gets the value for the given key.
		// `nil, nil` is returned when the key does not exist
		Get(key string) ([]byte, error)

		// Set stores the given value for the given key along with an expiration value, 0 means no expiration.
		// Empty key or value will be ignored without an error.
		Set(key string, val []byte, exp time.Duration) error

		// Delete deletes the value for the given key.
		// It returns no error if the storage does not contain the key,
		Delete(key string) error

		// Reset resets the storage and delete all keys.
		Reset() error

		// Close closes the storage and will stop any running garbage collectors and open connections.
		Close() error
	}
)
