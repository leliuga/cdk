package types

import (
	"fmt"
	"reflect"
	"slices"
)

// NewOption creates a new Option instance.
func NewOption(name String, description String, required bool, t Type, def any, min any, max any, choices []String) *Option {
	var chs []String
	for _, choice := range choices {
		chs = append(chs, choice)
	}
	return &Option{
		Name:        name,
		Description: description,
		Required:    required,
		Type:        &t,
		Default:     def,
		Min:         min,
		Max:         max,
		Choices:     chs,
	}
}

// NewBooleanOption creates a new Option instance with type boolean.
func NewBooleanOption(name String, description String, required bool, def bool) *Option {
	return NewOption(name, description, required, TypeBoolean, def, nil, nil, nil)
}

// NewDateTimeOption creates a new Option instance with type datetime.
func NewDateTimeOption(name String, description String, required bool, def String, min String, max String) *Option {
	return NewOption(name, description, required, TypeDateTime, def, min, max, nil)
}

// NewFloatOption creates a new Option instance with type float.
func NewFloatOption(name String, description String, required bool, def float64, min float64, max float64) *Option {
	return NewOption(name, description, required, TypeFloat, def, min, max, nil)
}

// NewIDOption creates a new Option instance with type id.
func NewIDOption(name String, description String, required bool, def String) *Option {
	return NewOption(name, description, required, TypeID, def, nil, nil, nil)
}

// NewIntegerOption creates a new Option instance with type integer.
func NewIntegerOption(name String, description String, required bool, def int, min int, max int) *Option {
	return NewOption(name, description, required, TypeInteger, def, min, max, nil)
}

// NewStringOption creates a new Option instance with type string.
func NewStringOption(name String, description String, required bool, def String, choices []String) *Option {
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
func (o Options) Index(name String) int {
	for index, option := range o {
		if option.Name == name {
			return index
		}
	}

	return -1
}

// Get returns the option with the given name.
func (o Options) Get(name String) (*Option, error) {
	index := o.Index(name)
	if index == -1 {
		return nil, fmt.Errorf("option %s not found", name)
	}

	return o[index], nil
}

// DefaultValue returns the default value of the option with the given name.
func (o Options) DefaultValue(name String, value any) any {
	if index := o.Index(name); index != -1 {
		return o[index].Default
	}

	return value
}

// SetDefaultValue sets the default value of the option with the given name.
func (o Options) SetDefaultValue(name String, value any) error {
	index := o.Index(name)
	if index == -1 {
		return fmt.Errorf("option %s not found", name)
	}

	kind := reflect.TypeOf(value).String()

	switch *o[index].Type {
	case TypeBoolean:
		if kind != "bool" {
			return fmt.Errorf("option %s is not of type boolean", name)
		}

		o[index].Default = value.(bool)
	case TypeDateTime:
		if kind != "string" {
			return fmt.Errorf("option %s is not of type datetime", name)
		}

		o[index].Default = value.(string)
	case TypeFloat:
		if !slices.Contains([]string{"float32", "float64"}, kind) {
			return fmt.Errorf("option %s is not of type float", name)
		}

		switch kind {
		case "float32":
			o[index].Default = value.(float32)
		default:
			o[index].Default = value.(float64)
		}
	case TypeID:
		if kind != "string" {
			return fmt.Errorf("option %s is not of type id", name)
		}

		o[index].Default = value.(string)
	case TypeInteger:
		if !slices.Contains([]string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"}, kind) {
			return fmt.Errorf("option %s is not of type integer", name)
		}

		switch kind {
		case "int":
			o[index].Default = value.(int)
		case "int8":
			o[index].Default = value.(int8)
		case "int16":
			o[index].Default = value.(int16)
		case "int32":
			o[index].Default = value.(int32)
		case "int64":
			o[index].Default = value.(int64)
		case "uint":
			o[index].Default = value.(uint)
		case "uint8":
			o[index].Default = value.(uint8)
		case "uint16":
			o[index].Default = value.(uint16)
		case "uint32":
			o[index].Default = value.(uint32)
		case "uint64":
			o[index].Default = value.(uint64)
		}
	case TypeString:
		if kind != "string" {
			return fmt.Errorf("option %s is not of type string", name)
		}

		o[index].Default = value.(string)
	}

	return nil
}

// ToSlice returns the options as a slice of strings.
func (o Options) ToSlice() []String {
	var arguments []String
	for _, option := range o {
		if option.Default == nil || option.Default == "" {
			continue
		}
		switch *option.Type {
		case TypeBoolean:
			if option.Default.(bool) {
				arguments = append(arguments, option.Name)
			}
		case TypeFloat:
			arguments = append(arguments, option.Name, Sprintf("%f", option.Default))
		case TypeInteger:
			arguments = append(arguments, option.Name, Sprintf("%d", option.Default))
		default:
			arguments = append(arguments, option.Name, Sprintf("%s", option.Default))
		}
	}

	return arguments
}
