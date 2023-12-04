package tokenizer

import (
	"regexp"
	"strings"

	"github.com/leliuga/cdk/types"
)

const (
	WordRegex = `[a-zą-žA-ZĄ-Ž]+`
)

var (
	SpaceRule        = NewRule("space", func(t types.String) bool { return t == " " })
	ExclamationRule  = NewRule("exclamation", func(t types.String) bool { return t == "!" })
	AtRule           = NewRule("at", func(t types.String) bool { return t == "@" })
	HashRule         = NewRule("hash", func(t types.String) bool { return t == "#" })
	DollarRule       = NewRule("dollar", func(t types.String) bool { return t == "$" })
	PercentRule      = NewRule("percent", func(t types.String) bool { return t == "%" })
	CaretRule        = NewRule("caret", func(t types.String) bool { return t == "^" })
	AmpersandRule    = NewRule("ampersand", func(t types.String) bool { return t == "&" })
	AsteriskRule     = NewRule("asterisk", func(t types.String) bool { return t == "*" })
	PipeRule         = NewRule("pipe", func(t types.String) bool { return t == "|" })
	DotRule          = NewRule("dot", func(t types.String) bool { return t == "." })
	CommaRule        = NewRule("comma", func(t types.String) bool { return t == "," })
	SemicolonRule    = NewRule("semicolon", func(t types.String) bool { return t == ";" })
	ColonRule        = NewRule("colon", func(t types.String) bool { return t == ":" })
	SingleQuoteRule  = NewRule("single_quote", func(t types.String) bool { return t == "'" })
	DoubleQuoteRule  = NewRule("double_quote", func(t types.String) bool { return t == "\"" })
	EqualRule        = NewRule("equal", func(t types.String) bool { return t == "=" })
	NotEqualRule     = NewRule("not_equal", func(t types.String) bool { return t == "!=" })
	LessRule         = NewRule("less", func(t types.String) bool { return t == "<" })
	LessEqualRule    = NewRule("less_equal", func(t types.String) bool { return t == "<=" })
	GreaterRule      = NewRule("greater", func(t types.String) bool { return t == ">" })
	GreaterEqualRule = NewRule("greater_equal", func(t types.String) bool { return t == ">=" })
	PlusRule         = NewRule("plus", func(t types.String) bool { return t == "+" })
	MinusRule        = NewRule("minus", func(t types.String) bool { return t == "-" })
	SlashRule        = NewRule("slash", func(t types.String) bool { return t == "/" })
	BackslashRule    = NewRule("backslash", func(t types.String) bool { return t == "\\" })
	QuestionRule     = NewRule("question", func(t types.String) bool { return t == "?" })
	WhitespaceRule   = NewRule("whitespace", func(t types.String) bool { return t == " " || t == "\t" || t == "\n" || t == "\r" })
	NewlineRule      = NewRule("newline", func(t types.String) bool { return t == "\n" || t == "\r" })
	LeftParenRule    = NewRule("left_paren", func(t types.String) bool { return t == "(" })
	RightParenRule   = NewRule("right_paren", func(t types.String) bool { return t == ")" })
	LeftBracketRule  = NewRule("left_bracket", func(t types.String) bool { return t == "[" })
	RightBracketRule = NewRule("right_bracket", func(t types.String) bool { return t == "]" })
	LeftBraceRule    = NewRule("left_brace", func(t types.String) bool { return t == "{" })
	RightBraceRule   = NewRule("right_brace", func(t types.String) bool { return t == "}" })
	UnderlineRule    = NewRule("underline", func(t types.String) bool { return t == "_" })
	FloatRule        = NewRule("float", func(t types.String) bool { _, err := t.ToFloat(64); return err == nil })
	IntegerRule      = NewRule("integer", func(t types.String) bool { _, err := t.ToInt(10, 0); return err == nil })
	LinkRule         = NewRule("link", func(t types.String) bool {
		u, err := t.ToURI()
		return err == nil && strings.Contains(u.Scheme, "http")
	})
	EmailRule = NewRule("email", func(t types.String) bool {
		u, err := t.ToURI()
		return err == nil && strings.Contains(u.Scheme, "mailto")
	})
	PhoneRule = NewRule("phone", func(t types.String) bool {
		u, err := t.ToURI()
		return err == nil && strings.Contains(u.Scheme, "tel")
	})
	WordRule = NewRule("word", func(t types.String) bool { return regexp.MustCompile(WordRegex).MatchString(t.String()) })
)

// NewRule creates a new Rule instance.
func NewRule(tag string, matcher Matcher) *Rule {
	return &Rule{
		Tag:     tag,
		Matcher: matcher,
	}
}
