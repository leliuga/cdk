package constants

import (
	"regexp"
)

var (
	NameRegex = regexp.MustCompile(`^[A-Z]+[a-z0-9 ]{0,62}$`)
)

// Default values
const (
	DefaultCharset        = "UTF-8"
	DefaultDateFormat     = "2006-01-02"
	DefaultDateTimeFormat = "2006-01-02 15:04:05"
	DefaultTimeFormat     = "15:04:05"
)

const (
	InvalidName = "A name must start with a capital letter and may contain only letters, numbers and space."
)
