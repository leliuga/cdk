package types

import (
	"errors"
	"io"
	"os"
	"path"
)

// UnmarshalFile unmarshals a file.
func UnmarshalFile(name string, out any) error {
	contentType := ContentTypeInvalid
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	switch path.Ext(name) {
	case ".yaml", ".yml":
		contentType = ContentTypeYaml
	case ".json":
		contentType = ContentTypeJson
	case ".msgpack":
		contentType = ContentTypeMsgPack
	default:
		return errors.New("invalid file extension")
	}

	return contentType.Unmarshal(f, out)
}

// MarshalFile marshals a file.
func MarshalFile(name string, in any) error {
	contentType := ContentTypeInvalid
	switch path.Ext(name) {
	case ".yaml", ".yml":
		contentType = ContentTypeYaml
	case ".json":
		contentType = ContentTypeJson
	case ".msgpack":
		contentType = ContentTypeMsgPack
	default:
		return errors.New("invalid file extension")
	}

	r, err := contentType.Marshal(in)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	return os.WriteFile(name, b, 0644)
}
