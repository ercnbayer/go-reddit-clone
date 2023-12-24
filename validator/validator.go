package validator

import (
	"emreddit/logger"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateID(id string) error {

	return Validate.Var(id, "uuid4") //validate id
}

func ValidateUpdatedStruct(user interface{}) error {

	logger.Info("validating no null fields only") // validate no null fields only

	return Validate.StructExcept(user)

}

func ValidateStruct(user interface{}) error {

	logger.Info("validating normal") //validate every field
	return Validate.Struct(user)
}
