package service

import (
	"bytes"
	"strings"
)

const (
	ProviderInvalid Provider = iota //
	ProviderAws
	ProviderAzure
	ProviderBareMetal
	ProviderDo
	ProviderGcp
)

var (
	ProviderNames = map[Provider]string{
		ProviderAws:       "Amazon Web Service",
		ProviderAzure:     "Azure",
		ProviderBareMetal: "Bare Metal",
		ProviderDo:        "Digital Ocean",
		ProviderGcp:       "Google Cloud Platform",
	}
)

// String outputs the Provider as a string.
func (p Provider) String() string {
	return ProviderNames[p]
}

// MarshalJSON outputs the Provider as a json.
func (p Provider) MarshalJSON() ([]byte, error) {
	if !p.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + p.String() + `"`), nil
}

// UnmarshalJSON parses the Provider from json.
func (p *Provider) UnmarshalJSON(data []byte) error {
	str := string(bytes.Trim(data, `"`))
	if provider := ParseProvider(str); provider.Validate() {
		*p = provider
	}

	return nil
}

// Validate returns true if the Provider is valid.
func (p Provider) Validate() bool {
	return p != ProviderInvalid
}

// ParseProvider parses the Provider from string.
func ParseProvider(value string) Provider {
	value = strings.ToLower(value)
	for k, v := range ProviderNames {
		if strings.ToLower(v) == v {
			return k
		}
	}

	return ProviderInvalid
}
