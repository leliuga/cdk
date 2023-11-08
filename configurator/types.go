package configurator

type (
	// parserFunc is a function that parses a string into a value.
	parserFunc func(value string) (any, error)

	// EnvironmentEntry is a key-value pair that represents an environment variable.
	EnvironmentEntry struct {
		Key    string `json:"key"`
		Value  string `json:"value"`
		Secret bool   `json:"secret"`
	}

	// Environment is a list of environment variables.
	Environment struct {
		Name        string              `json:"name"`
		Description string              `json:"description"`
		Entries     []*EnvironmentEntry `json:"entries"`
	}
)
