package utils

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func DecodeAndValidate(reader io.ReadCloser, obj any) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	return Validator.Struct(obj)
}

func PrettifyJSON(obj any) (string, error) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return "```json\n" + string(data) + "\n```", nil
}
