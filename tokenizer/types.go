package tokenizer

import (
	"github.com/leliuga/cdk/types"
)

type (
	// Tokenizer represents the tokenizer.
	Tokenizer struct {
		*Options `json:",inline"`

		cache map[types.String]*Token
	}

	// Options represents the tokenizer options.
	Options struct {
		Splitter Splitter `json:"splitter"`
		Rules    Rules    `json:"rules"`
	}

	// Rule is a rule
	Rule struct {
		Tag     string  `json:"tag"`
		Matcher Matcher `json:"matcher"`
	}

	// Token represents the token.
	Token struct {
		Tag   string       `json:"tag"`
		Value types.String `json:"value"`
	}

	// Rules represents the rules.
	Rules []*Rule

	// Tokens represents the tokens.
	Tokens []*Token

	// Option represents the tokenizer option.
	Option func(*Options)

	// Splitter represents the extractor.
	Splitter func(types.String) []types.String

	// Matcher represents the matcher.
	Matcher func(types.String) bool
)
