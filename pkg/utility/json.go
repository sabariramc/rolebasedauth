package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonTransformer(src interface{}, dest interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(src)
	if err != nil {
		return fmt.Errorf("JsonTransformer encoding: %w", err)
	}
	err = json.NewDecoder(&buf).Decode(dest)
	if err != nil {
		return fmt.Errorf("JsonTransformer decoding: %w", err)
	}
	return nil
}
