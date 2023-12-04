package tokenizer

import (
	"github.com/leliuga/cdk/types"
)

const (
	// DefaultPattern is the default pattern.
	DefaultPattern = `(?:\p{L}+|\p{N}+(\.\p{N}+)?|>=|<=|!=|\ |[^\p{L}\p{N}\s])`
)

// NewOptions creates a new Options instance.
func NewOptions(options ...Option) *Options {
	opts := Options{
		Splitter: func(text types.String) []types.String {
			return text.FindAllString(DefaultPattern, -1)
		},
		Rules: Rules{
			SpaceRule,
			ExclamationRule,
			AtRule,
			HashRule,
			DollarRule,
			PercentRule,
			CaretRule,
			AmpersandRule,
			AsteriskRule,
			PipeRule,
			DotRule,
			CommaRule,
			SemicolonRule,
			ColonRule,
			SingleQuoteRule,
			DoubleQuoteRule,
			EqualRule,
			NotEqualRule,
			LessRule,
			LessEqualRule,
			GreaterRule,
			GreaterEqualRule,
			PlusRule,
			MinusRule,
			SlashRule,
			BackslashRule,
			QuestionRule,
			WhitespaceRule,
			NewlineRule,
			LeftParenRule,
			RightParenRule,
			LeftBracketRule,
			RightBracketRule,
			LeftBraceRule,
			RightBraceRule,
			UnderlineRule,
			IntegerRule,
			FloatRule,
			LinkRule,
			EmailRule,
			PhoneRule,
			WordRule,
		},
	}

	for _, option := range options {
		option(&opts)
	}

	return &opts
}

// WithSplitter sets the splitter.
func WithSplitter(splitter Splitter) Option {
	return func(options *Options) {
		options.Splitter = splitter
	}
}

// WithRules sets the rules.
func WithRules(rules Rules) Option {
	return func(options *Options) {
		options.Rules = rules
	}
}
