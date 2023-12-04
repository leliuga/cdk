package tokenizer

func (t Tokens) String() string {
	var s string

	for k, token := range t {
		s += token.String()

		if k != len(t)-1 {
			s += "\n"
		}
	}
	return s
}

// Is checks if the tokens are of the given tags.
func (t Tokens) Is(tag string) bool {
	for _, token := range t {
		if !token.Is(tag) {
			return false
		}
	}

	return true
}

// Only returns the tokens of the given tag.
func (t Tokens) Only(tag string) Tokens {
	var tokens Tokens
	for _, token := range t {
		if token.Is(tag) {
			tokens = append(tokens, token)
		}
	}

	return tokens
}
