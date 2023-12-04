package tokenizer

import (
	"github.com/leliuga/cdk/types"
)

// NewToken creates a new Token instance.
func NewToken(value types.String, rules Rules) *Token {
	tag := ""
	for _, rule := range rules {
		if rule.Matcher(value.ToLower()) {
			tag = rule.Tag
			break
		}
	}

	return &Token{
		Tag:   tag,
		Value: value,
	}
}

// Is checks if the token is of the given tag.
func (t *Token) Is(tag string) bool {
	return t.Tag == tag
}

// String returns the token as a string.
func (t *Token) String() string {
	return types.Sprintf("v: %s t: %s", t.Value, t.Tag).String()
}
