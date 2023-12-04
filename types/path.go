package types

import (
	"bytes"
	"database/sql/driver"

	"github.com/leliuga/cdk/tokenizer"
)

// String returns the path as a string.
func (p *Path) String() string {
	return p.p.String()
}

// Bytes returns the path as bytes.
func (p *Path) Bytes() []byte {
	return p.p.Bytes()
}

// Value outputs the URI as a value.
func (p *Path) Value() (driver.Value, error) {
	return p.String(), nil
}

// MarshalJSON outputs the URI as a json.
func (p *Path) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}

// UnmarshalJSON parses URI from json.
func (p *Path) UnmarshalJSON(data []byte) error {
	v, err := ParsePath(String(bytes.Trim(data, `"`)))
	if err != nil {
		return nil
	}

	*p = *v

	return nil
}

// Hash returns the path as a hash.
func (p *Path) Hash() uint32 {
	return p.hash
}

// Segments returns the path as a list of segments.
func (p *Path) Segments() []String {
	return p.p.Segments("/")
}

// Equal checks if two paths are equal.
func (p *Path) Equal(p1 *Path) bool {
	return p.hash == p1.hash
}

// Tokenize tokenizes the path.
func (p *Path) Tokenize() tokenizer.Tokens {
	return p.p.Tokenize(tokenizer.NewOptions(
		tokenizer.WithSplitter(func(text String) []String {
			return text.Trim("/").Segments("/")
		}),
		tokenizer.WithRules(tokenizer.Rules{
			tokenizer.NewRule("strict", func(t String) bool { return !t.HasPrefix(":") && !t.HasPrefix("*") }),
			tokenizer.NewRule("param", func(t String) bool { return t.HasPrefix(":") }),
			tokenizer.NewRule("wildcard", func(t String) bool { return t.HasPrefix("*") }),
		}),
	))
}

// ParsePath parses Path from string.
func ParsePath(value String) (*Path, error) {
	path := value.TrimSpace()

	return &Path{
		p:    path,
		hash: path.Hash(),
	}, nil
}
