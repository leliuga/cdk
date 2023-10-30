package resource

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	KindInvalid Kind = iota
	KindCPU
	KindMemory
	KindSwap
	KindStorage
	KindEphemeralStorage
	KindNetwork
	KindGPU
	KindTPU
)

var (
	KindNames = map[Kind]string{
		KindCPU:              "cpu",
		KindMemory:           "memory",
		KindSwap:             "swap",
		KindStorage:          "storage",
		KindEphemeralStorage: "ephemeral-storage",
		KindNetwork:          "network",
		KindGPU:              "gpu",
		KindTPU:              "tpu",
	}
)

// String outputs the Kind as a string.
func (k Kind) String() string {
	return KindNames[k]
}

// Bytes returns the Kind as a []byte.
func (k Kind) Bytes() []byte {
	return []byte(strconv.Itoa(int(k)))
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
		if v == value {
			return k
		}
	}

	return KindInvalid
}
