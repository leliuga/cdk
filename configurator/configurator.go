package configurator

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
	"github.com/leliuga/cdk/types"
	"k8s.io/apimachinery/pkg/api/resource"
)

// FromEnv reads the environment variables and sets the values to the given config.
func FromEnv(config any, prefix string) {
	structureIteration(config, func(fieldStruct reflect.StructField, field reflect.Value, envName string) {
		envValue := os.Getenv(prefix + envName)

		if envValue == "" && field.Kind() != reflect.Struct {
			return
		}

		switch field.Kind() {
		case reflect.Struct:
			FromEnv(field.Addr().Interface(), prefix+envName+"_")
		case reflect.Slice:
			values := strings.Split(envValue, ",")
			slice := reflect.MakeSlice(field.Type(), len(values), len(values))
			for key, value := range values {
				setFieldValue(slice.Index(key), value)
			}
			field.Set(slice)
		case reflect.Map:
			m := reflect.MakeMap(field.Type())
			for _, pair := range strings.Split(envValue, ",") {
				kv := strings.Split(pair, "=")
				if len(kv) != 2 {
					continue
				}

				valueType := fieldStruct.Type.Elem()
				v, err := typeParser(valueType, kv[1])
				if err != nil {
					panic(err)
				}

				keyType := fieldStruct.Type.Key()
				m.SetMapIndex(
					reflect.ValueOf(kv[0]).Convert(keyType),
					reflect.ValueOf(v).Convert(valueType),
				)
			}
			field.Set(m)
		default:
			setFieldValue(field, envValue)
		}
	})
}

// ToEnv returns an environment struct with the values of the given config.
func ToEnv(config any, name, description, prefix string) *Environment {
	env := NewEnvironment(name, description)

	structureIteration(config, func(fieldStruct reflect.StructField, field reflect.Value, envName string) {
		switch field.Kind() {
		case reflect.Struct:
			for _, e := range ToEnv(field.Addr().Interface(), name, description, prefix+envName+"_").Entries {
				env.Set(&EnvironmentEntry{
					Key:   e.Key,
					Value: e.Value,
				})
			}
		case reflect.Slice:
			value := reflect.ValueOf(field.Interface())
			values := make([]string, 0)
			for i := 0; i < value.Len(); i++ {
				val, _ := json.Marshal(value.Index(i).Interface())
				values = append(values, fmt.Sprintf("%s", val))
			}

			env.Set(&EnvironmentEntry{
				Key:   prefix + envName,
				Value: strings.Join(values, ","),
			})
		case reflect.Map:
			iter := reflect.ValueOf(field.Interface()).MapRange()
			values := make([]string, 0)
			for iter.Next() {
				key := iter.Key().Interface()
				val, _ := json.Marshal(iter.Value().Interface())

				values = append(values, fmt.Sprintf("%s=%s", key, val))
			}

			env.Set(&EnvironmentEntry{
				Key:   prefix + envName,
				Value: strings.Join(values, ","),
			})
		default:
			env.Set(&EnvironmentEntry{
				Key:   prefix + envName,
				Value: fmt.Sprintf("%v", field.Interface()),
			})
		}
	})

	return env
}

// structureIteration iterates over the fields of a struct and calls the given function for each field.
func structureIteration(config any, fn func(reflect.StructField, reflect.Value, string)) {
	if config == nil {
		return
	}

	configValue := reflect.Indirect(reflect.ValueOf(config))
	if configValue.Kind() != reflect.Struct {
		return
	}

	configType := configValue.Type()
	for index := 0; index < configType.NumField(); index++ {
		fieldStruct := configType.Field(index)
		field := configValue.Field(index)
		envName := fieldStruct.Tag.Get("env")

		if envName == "" || !field.CanAddr() || !field.CanInterface() {
			continue
		}

		// Traverse pointers and interfaces to get the actual field value
		for field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface {
			if field.IsNil() {
				break
			}
			field = field.Elem()
		}

		fn(fieldStruct, field, envName)
	}
}

// setFieldValue sets the value of a field to the given value.
func setFieldValue(field reflect.Value, value string) {
	v, err := kindParser(field.Kind(), value)
	if err != nil {
		panic(err)
	}

	field.Set(reflect.ValueOf(v).Convert(field.Type()))
}

// kindParser returns a function that parses a string to the given kind.
func kindParser(k reflect.Kind, value string) (any, error) {
	parsers := map[reflect.Kind]parserFunc{
		reflect.String: func(v string) (any, error) {
			return v, nil
		},
		reflect.Int: func(v string) (any, error) {
			i, err := strconv.ParseInt(v, 10, 32)
			return int(i), err
		},
		reflect.Int8: func(v string) (any, error) {
			i, err := strconv.ParseInt(v, 10, 8)
			return int8(i), err
		},
		reflect.Int16: func(v string) (any, error) {
			i, err := strconv.ParseInt(v, 10, 16)
			return int16(i), err
		},
		reflect.Int32: func(v string) (any, error) {
			i, err := strconv.ParseInt(v, 10, 32)
			return int32(i), err
		},
		reflect.Int64: func(v string) (any, error) {
			return strconv.ParseInt(v, 10, 64)
		},
		reflect.Uint: func(v string) (any, error) {
			i, err := strconv.ParseUint(v, 10, 32)
			return uint(i), err
		},
		reflect.Uint8: func(v string) (any, error) {
			i, err := strconv.ParseUint(v, 10, 8)
			return uint8(i), err
		},
		reflect.Uint16: func(v string) (any, error) {
			i, err := strconv.ParseUint(v, 10, 16)
			return uint16(i), err
		},
		reflect.Uint32: func(v string) (any, error) {
			i, err := strconv.ParseUint(v, 10, 32)
			return uint32(i), err
		},
		reflect.Uint64: func(v string) (any, error) {
			return strconv.ParseUint(v, 10, 64)
		},
		reflect.Float32: func(v string) (any, error) {
			f, err := strconv.ParseFloat(v, 32)
			return float32(f), err
		},
		reflect.Float64: func(v string) (any, error) {
			return strconv.ParseFloat(v, 64)
		},
		reflect.Bool: func(v string) (any, error) {
			return strconv.ParseBool(v)
		},
	}

	parseFunc, ok := parsers[k]
	if !ok {
		return nil, fmt.Errorf("no parser for kind %s", k)
	}

	return parseFunc(value)
}

// typeParser returns a function that parses a string to the given type.
func typeParser(t reflect.Type, value string) (any, error) {
	parsers := map[reflect.Type]parserFunc{
		reflect.TypeOf(types.URI{}): func(v string) (any, error) {
			return types.ParseURI(v), nil
		},
		reflect.TypeOf(resource.Quantity{}): func(v string) (any, error) {
			return resource.ParseQuantity(v)
		},
	}

	parseFunc, ok := parsers[t]
	if !ok {
		return nil, fmt.Errorf("no parser for type %s", t)
	}

	return parseFunc(value)
}
