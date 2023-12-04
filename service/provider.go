package service

import (
	"bytes"
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
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
func (p *Provider) String() string {
	if !p.Validate() {
		return ""
	}

	return ProviderNames[*p]
}

// Bytes returns the Provider as a []byte.
func (p *Provider) Bytes() []byte {
	return []byte(p.String())
}

// Value outputs the Provider as a value.
func (p *Provider) Value() (driver.Value, error) {
	return p.String(), nil
}

// MarshalJSON outputs the Provider as a json.
func (p *Provider) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}

// UnmarshalJSON parses the Provider from json.
func (p *Provider) UnmarshalJSON(data []byte) error {
	v, err := ParseProvider(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*p = *v

	return nil
}

// Validate returns true if the Provider is valid.
func (p *Provider) Validate() bool {
	return *p != ProviderInvalid
}

// ParseProvider parses the Provider from string.
func ParseProvider(value string) (*Provider, error) {
	value = strings.ToLower(value)
	for k, v := range ProviderNames {
		if strings.ToLower(v) == value {
			return &k, nil
		}
	}

	return nil, errors.Errorf("unsupported provider: %s", value)
}
