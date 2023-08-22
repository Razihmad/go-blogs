package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct(v interface{}) error {
	err := validate.Struct(v)
	if err != nil {
		return err
	}
	return nil
}
