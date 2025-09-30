package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func EncodeStruct(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct: %w", err)
	}
	return base64.URLEncoding.EncodeToString(jsonData), nil
}
