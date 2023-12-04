package event

import (
	"bufio"
	"net/url"

	"github.com/google/uuid"
	"github.com/leliuga/cdk/types"
)

// NewOptions returns new options.
func NewOptions(options ...Option) *Options {
	opts := Options{
		ID:         uuid.New().String(),
		Attributes: types.NewMap[string](),
		Data:       []byte{},
		Happen:     types.DateTimeNow(),
	}

	for _, option := range options {
		option(&opts)
	}

	return &opts
}

// WithID sets the id for the event.
func WithID(value string) Option {
	return func(o *Options) {
		o.ID = value
	}
}

// WithVersion sets the version for the event.
func WithVersion(value string) Option {
	return func(o *Options) {
		o.Version = value
	}
}

// WithSchema sets the schema for the event.
func WithSchema(value string) Option {
	return func(o *Options) {
		o.Schema, _ = types.ParseURI(value)
	}
}

// WithSource sets the source for the event.
func WithSource(value string) Option {
	return func(o *Options) {
		o.Source, _ = types.ParseURI(value)
	}
}

// WithKind sets the kind for the event.
func WithKind(value Kind) Option {
	return func(o *Options) {
		o.Kind = value
	}
}

// WithAction sets the action for the event.
func WithAction(value Action) Option {
	return func(o *Options) {
		o.Action = value
	}
}

// WithAttributes sets the attributes for the event.
func WithAttributes(value types.Map[string]) Option {
	return func(o *Options) {
		o.Attributes = value
	}
}

// WithData sets the data for the event.
func WithData(value []byte) Option {
	return func(o *Options) {
		o.Data = value
	}
}

// WithJsonData sets the json data for the event.
func WithJsonData(value any) Option {
	return func(o *Options) {
		json := types.ContentTypeJson
		reader, _ := json.Marshal(value)
		buf := bufio.NewScanner(reader)
		o.Data, _ = buf.Bytes(), nil
	}
}

// WithMsgPackData sets the msgpack data for the event.
func WithMsgPackData(value any) Option {
	return func(o *Options) {
		o.Data, _ = types.ContentTypeMsgPack.Marshal(value)
	}
}

// WithYamlData sets the yaml data for the event.
func WithYamlData(value any) Option {
	return func(o *Options) {
		o.Data, _ = types.ContentTypeYaml.Marshal(value)
	}
}

// WithFormUrlEncodedData sets the form url encoded data for the event.
func WithFormUrlEncodedData(value url.Values) Option {
	return func(o *Options) {
		o.Data, _ = types.ContentTypeFormUrlEncoded.Marshal(value)
	}
}

// WithHappen sets the happen for the event.
func WithHappen(value *types.DateTime) Option {
	return func(o *Options) {
		o.Happen = value
	}
}
