package tokenizer

import (
	"github.com/leliuga/cdk/types"
)

// NewTokenizer creates a new Tokenizer instance.
func NewTokenizer(options *Options) *Tokenizer {
	return &Tokenizer{
		Options: options,
		cache:   map[types.String]*Token{},
	}
}

// Tokenize tokenizes the given text.
func (t *Tokenizer) Tokenize(text types.String) Tokens {
	var tokens Tokens
	values := t.Splitter(text)

	if len(values) == 0 {
		return tokens
	}

	for _, value := range values {
		if token, ok := t.cache[value]; ok {
			tokens = append(tokens, token)
			continue
		}

		t.cache[value] = NewToken(value, t.Rules)
		tokens = append(tokens, t.cache[value])
	}

	return tokens
}
