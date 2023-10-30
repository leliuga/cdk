package event

import (
	"bytes"
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
func (k Kind) String() string {
	return KindNames[k]
}

// MarshalJSON outputs the Kind as a json.
func (k Kind) MarshalJSON() ([]byte, error) {
	if !k.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + k.String() + `"`), nil
}

// UnmarshalJSON parses Kind from json.
func (k *Kind) UnmarshalJSON(data []byte) error {
	s := string(bytes.Trim(data, `"`))
	if r := ParseKind(s); r.Validate() {
		*k = r
	}

	return nil
}

// Validate returns true if the Kind is valid.
func (k Kind) Validate() bool {
	return k != KindInvalid
}

// ParseKind parses Kind from string.
func ParseKind(value string) Kind {
	value = strings.ToLower(value)
	for k, v := range KindNames {
		if strings.ToLower(v) == value {
			return k
		}
	}

	return KindInvalid
}
