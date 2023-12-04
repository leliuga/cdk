package types

import (
	"database/sql/driver"
	"fmt"
	"hash/fnv"
	"regexp"
	"strconv"
	"strings"

	"github.com/leliuga/cdk/tokenizer"
)

// NewString returns a new string.
func NewString(value string) String {
	return String(value)
}

// Sprintf returns a new string.
func Sprintf(format string, a ...any) String {
	return String(fmt.Sprintf(format, a...))
}

// FromBytes returns a new string.
func FromBytes(b []byte) String {
	return String(b)
}

// FromInt returns a new string.
func FromInt(i int64, base int) String {
	return String(strconv.FormatInt(i, base))
}

// FromUint returns a new string.
func FromUint(i uint64, base int) String {
	return String(strconv.FormatUint(i, base))
}

// FromFloat returns a new string.
func FromFloat(f float64, fmt byte, prec int, bitSize int) String {
	return String(strconv.FormatFloat(f, fmt, prec, bitSize))
}

// FromBool returns a new string.
func FromBool(b bool) String {
	return String(strconv.FormatBool(b))
}

// String returns the path as a string.
func (s String) String() string {
	return string(s)
}

// Bytes returns the path as bytes.
func (s String) Bytes() []byte {
	return []byte(s)
}

// Value outputs the URI as a value.
func (s String) Value() (driver.Value, error) {
	return s, nil
}

// MarshalJSON outputs the URI as a json.
func (s String) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s + `"`), nil
}

// UnmarshalJSON parses URI from json.
func (s *String) UnmarshalJSON(data []byte) error {
	*s = String(data)

	return nil
}

// Validate returns true if the URI is valid.
func (s String) Validate() bool {
	return true
}

// Contains checks if a string contains another string.
func (s String) Contains(substr String) bool {
	return strings.Contains(string(s), string(substr))
}

// Equal checks if two paths are equal.
func (s String) Equal(s1 String) bool {
	return s == s1
}

func (s String) IsEmpty() bool {
	return s == ""
}

// Len returns the length of the string.
func (s String) Len() int {
	return len(s)
}

// ToLower returns the string in lower case.
func (s String) ToLower() String {
	return String(strings.ToLower(string(s)))
}

// ToUpper returns the string in upper case.
func (s String) ToUpper() String {
	return String(strings.ToUpper(string(s)))
}

// ToTitle returns the string in title case.
func (s String) ToTitle() String {
	return String(strings.Title(string(s)))
}

// ToInt returns the string as an integer.
func (s String) ToInt(base int, bitSize int) (int64, error) {
	return strconv.ParseInt(string(s), base, bitSize)
}

// ToUint returns the string as an unsigned integer.
func (s String) ToUint(base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(string(s), base, bitSize)
}

// ToFloat returns the string as a float.
func (s String) ToFloat(bitSize int) (float64, error) {
	return strconv.ParseFloat(string(s), bitSize)
}

// ToBool returns the string as a boolean.
func (s String) ToBool() (bool, error) {
	return strconv.ParseBool(string(s))
}

// ToPath returns the string as a path.
func (s String) ToPath() (*Path, error) {
	return ParsePath(s)
}

// ToURI returns the string as a URI.
func (s String) ToURI() (*URI, error) {
	return ParseURI(string(s))
}

// Tokenize returns the string as a list of tokens.
func (s String) Tokenize(options *tokenizer.Options) tokenizer.Tokens {
	return tokenizer.NewTokenizer(options).Tokenize(s)
}

// TrimSpace returns the string with leading and trailing white space removed.
func (s String) TrimSpace() String {
	return String(strings.TrimSpace(string(s)))
}

// Trim returns the string with leading and trailing characters removed.
func (s String) Trim(cutset String) String {
	return String(strings.Trim(string(s), string(cutset)))
}

// TrimLeft returns the string with leading characters removed.
func (s String) TrimLeft(cutset String) String {
	return String(strings.TrimLeft(string(s), string(cutset)))
}

// TrimRight returns the string with trailing characters removed.
func (s String) TrimRight(cutset String) String {
	return String(strings.TrimRight(string(s), string(cutset)))
}

// TrimPrefix returns the string without the provided leading prefix string.
func (s String) TrimPrefix(prefix String) String {
	if !s.HasPrefix(prefix) {
		return s
	}

	return s[len(prefix):]
}

// TrimSuffix returns the string without the provided trailing suffix string.
func (s String) TrimSuffix(suffix String) String {
	if !s.HasSuffix(suffix) {
		return s
	}

	return s[:len(s)-len(suffix)]
}

// HasPrefix returns true if the string starts with the provided prefix.
func (s String) HasPrefix(prefix String) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// HasSuffix returns true if the string ends with the provided suffix.
func (s String) HasSuffix(suffix String) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Segments returns the string as a list of segments.
func (s String) Segments(sep String) (out Slice[String]) {
	for _, value := range strings.Split(string(s.Trim(sep)), string(sep)) {
		out = out.Append(String(value))
	}

	return out
}

// Join returns the string joined by the separator.
func (s String) Join(elems Slice[String], sep String) (out String) {
	if elems.Len() == 0 {
		return ""
	}

	for k, v := range elems {
		out += v

		if k != elems.Len()-1 {
			out += sep
		}
	}

	return out
}

// Replace returns the string with the old replaced by the new.
func (s String) Replace(old, new String) String {
	return String(strings.Replace(string(s), string(old), string(new), -1))
}

// ReplaceAll returns the string with the old replaced by the new.
func (s String) ReplaceAll(old, new String) String {
	return String(strings.ReplaceAll(string(s), string(old), string(new)))
}

// FindAllString returns the string with the old replaced by the new.
func (s String) FindAllString(pattern String, n int) (out Slice[String]) {
	for _, value := range regexp.MustCompile(string(pattern)).FindAllString(string(s), n) {
		out = out.Append(String(value))
	}

	return out
}

// MatchString returns the string with the old replaced by the new.
func (s String) MatchString(pattern String) bool {
	return regexp.MustCompile(string(pattern)).MatchString(string(s))
}

// Hash returns the string as a hash.
func (s String) Hash() uint32 {
	h := fnv.New32a()
	_, _ = h.Write(s.Bytes())

	return h.Sum32()
}
