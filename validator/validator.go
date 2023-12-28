package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateUUID(id string) error {

	return Validate.Var(id, "uuid4") //validate id
}
