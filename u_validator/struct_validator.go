package u_validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Struct(v interface{}) error {
	return validate.Struct(v)
}
