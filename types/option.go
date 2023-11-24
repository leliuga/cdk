package types

// NewOption creates a new Option instance.
func NewOption(name string, description string, required bool, t Type, def any, min any, max any, choices []string) *Option {
	return &Option{
		Name:        name,
		Description: description,
		Required:    required,
		Type:        t,
		Default:     def,
		Min:         min,
		Max:         max,
		Choices:     choices,
	}
}

// NewBooleanOption creates a new Option instance with type boolean.
func NewBooleanOption(name string, description string, required bool, def bool) *Option {
	return NewOption(name, description, required, TypeBoolean, def, nil, nil, nil)
}

// NewFloatOption creates a new Option instance with type float.
func NewFloatOption(name string, description string, required bool, def float64, min float64, max float64) *Option {
	return NewOption(name, description, required, TypeFloat, def, min, max, nil)
}

// NewIntegerOption creates a new Option instance with type integer.
func NewIntegerOption(name string, description string, required bool, def int, min int, max int) *Option {
	return NewOption(name, description, required, TypeInteger, def, min, max, nil)
}

// NewStringOption creates a new Option instance with type string.
func NewStringOption(name string, description string, required bool, def string, choices []string) *Option {
	return NewOption(name, description, required, TypeString, def, nil, nil, choices)
}

// NewOptions creates a new Options instance.
func NewOptions(options ...*Option) Options {
	opts := Options{}
	for _, option := range options {
		opts = append(opts, option)
	}

	return opts
}

// Len returns the length of the options.
func (o Options) Len() int {
	return len(o)
}

// Index returns the index of the option with the given name.
func (o Options) Index(name string) int {
	for index, option := range o {
		if option.Name == name {
			return index
		}
	}

	return -1
}
