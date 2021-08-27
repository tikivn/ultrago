package validators

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Struct(v interface{}) error {
	return validate.Struct(v)
}
