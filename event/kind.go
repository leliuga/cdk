package event

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"
)

const (
	KindInvalid Kind = iota //
	KindApplicationTrace
	KindApplicationLog
	KindApplicationMetric
	KindKubernetesManifest
	KindKubernetesLog
	KindKubernetesMetric
)

var (
	KindNames = map[Kind]string{
		KindApplicationTrace:   "ApplicationTrace",
		KindApplicationLog:     "ApplicationLog",
		KindApplicationMetric:  "ApplicationMetric",
		KindKubernetesManifest: "KubernetesManifest",
		KindKubernetesLog:      "KubernetesLog",
		KindKubernetesMetric:   "KubernetesMetric",
	}
)

// String outputs the Kind as a string.
func (k *Kind) String() string {
	if !k.Validate() {
		return ""
	}

	return KindNames[*k]
}

// Bytes outputs the Kind as bytes.
func (k *Kind) Bytes() []byte {
	return []byte(k.String())
}

// Value outputs the Kind as a value.
func (k *Kind) Value() (driver.Value, error) {
	return k.String(), nil
}

// MarshalJSON outputs the Kind as a json.
func (k *Kind) MarshalJSON() ([]byte, error) {
	return []byte(`"` + k.String() + `"`), nil
}

// UnmarshalJSON parses Kind from json.
func (k *Kind) UnmarshalJSON(data []byte) error {
	v, err := ParseKind(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*k = *v

	return nil
}

// Validate returns true if the Kind is valid.
func (k *Kind) Validate() bool {
	return *k != KindInvalid
}

// ParseKind parses Kind from string.
func ParseKind(value string) (*Kind, error) {
	value = strings.ToLower(value)
	for k, v := range KindNames {
		if strings.ToLower(v) == value {
			return &k, nil
		}
	}

	return nil, fmt.Errorf("unsupported kind: %s", value)
}
