package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	ContentTypeInvalid ContentType = iota //
	ContentTypeJson
	ContentTypeMsgPack
	ContentTypeYaml
	ContentTypeHtml
	ContentTypeFormUrlEncoded
	ContentTypePlain
)

var (
	ContentTypeNames = map[ContentType]string{
		ContentTypeJson:           "application/json",
		ContentTypeMsgPack:        "application/msgpack",
		ContentTypeYaml:           "application/yaml",
		ContentTypeHtml:           "text/html",
		ContentTypeFormUrlEncoded: "application/x-www-form-urlencoded",
		ContentTypePlain:          "text/plain",
	}
)

// ContentTypeMarshal parses ContentType and returns io.Reader
func ContentTypeMarshal(contentType string, in any) (io.Reader, error) {
	ct, err := ParseContentType(contentType)
	if err != nil {
		return nil, err
	}

	return ct.Marshal(in)
}

// ContentTypeUnmarshal
func ContentTypeUnmarshal(contentType string, r io.Reader, out any) error {
	ct, err := ParseContentType(contentType)
	if err != nil {
		return err
	}

	return ct.Unmarshal(r, out)
}

// String outputs the ContentType as a string.
func (ct *ContentType) String() string {
	if !ct.Validate() {
		return ""
	}

	return ContentTypeNames[*ct]
}

// Bytes returns the ContentType as a []byte.
func (ct *ContentType) Bytes() []byte {
	return []byte(ct.String())
}

// Value outputs the ContentType as a value.
func (ct *ContentType) Value() (driver.Value, error) {
	return ct.String(), nil
}

// MarshalJSON outputs the ContentType as a json.
func (ct *ContentType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.String() + `"`), nil
}

// UnmarshalJSON parses ContentType from json.
func (ct *ContentType) UnmarshalJSON(data []byte) error {
	v, err := ParseContentType(string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	*ct = *v

	return nil
}

// Marshal returns the ContentType encoding of in.
func (ct *ContentType) Marshal(in any) (io.Reader, error) {
	buffer := bytes.NewBuffer(nil)
	switch *ct {
	case ContentTypeJson:
		if err := json.NewEncoder(buffer).Encode(in); err != nil {
			return nil, err
		}
	case ContentTypeMsgPack:
		if err := msgpack.NewEncoder(buffer).Encode(in); err != nil {
			return nil, err
		}
	case ContentTypeYaml:
		if err := yaml.NewEncoder(buffer, yaml.UseJSONMarshaler()).Encode(in); err != nil {
			return nil, err
		}
	case ContentTypeHtml, ContentTypePlain:
		buffer.WriteString(fmt.Sprintf("%v", in))
	case ContentTypeFormUrlEncoded:
		switch in.(type) {
		case string:
			buffer.WriteString(in.(string))
		case url.Values:
			buffer.WriteString(in.(url.Values).Encode())
		default:
			return nil, fmt.Errorf("a data type %T is invalid", in)
		}
	}

	return buffer, nil
}

// Unmarshal parses the ContentType-encoded data and stores the result in the value pointed to by out.
func (ct *ContentType) Unmarshal(r io.Reader, out any) error {
	switch *ct {
	case ContentTypeJson:
		return json.NewDecoder(r).Decode(out)
	case ContentTypeMsgPack:
		return msgpack.NewDecoder(r).Decode(out)
	case ContentTypeYaml:
		return yaml.NewDecoder(r, yaml.UseJSONUnmarshaler()).Decode(out)
	case ContentTypeHtml, ContentTypePlain, ContentTypeFormUrlEncoded:
		b, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		if *ct == ContentTypeFormUrlEncoded {
			out, err = url.ParseQuery(string(b))
			return err
		}

		out = string(b)
	}

	return nil
}

// Validate returns true if the ContentType is valid.
func (ct *ContentType) Validate() bool {
	return *ct != ContentTypeInvalid
}

// ParseContentType parses ContentType from string.
func ParseContentType(value string) (*ContentType, error) {
	parts := strings.Split(strings.ToLower(value), ";")
	value = parts[0]

	for k, v := range ContentTypeNames {
		if v == value {
			return &k, nil
		}
	}

	return nil, fmt.Errorf("unsupported content type: %s", value)
}
