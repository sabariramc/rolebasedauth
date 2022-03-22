package app

import (
	"bytes"
	"encoding/json"
)

func JsonTransformer(src interface{}, dest interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(src)
	if err != nil {
		return err
	}
	err = json.NewDecoder(&buf).Decode(dest)
	if err != nil {
		return err
	}
	return nil
}
