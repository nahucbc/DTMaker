package utils

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

func CustomEncoder(v any) ([]byte, error) {
	buff := new(bytes.Buffer)
	encoder := toml.NewEncoder(buff)
	encoder.Indent = ""
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
