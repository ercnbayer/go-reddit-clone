package validator

import (
	"emreddit/logger"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateID(id string) error {

	return Validate.Var(id, "uuid4")
}

func ValidateUpdatedStruct(user interface{}) error {

	logger.Info("validating partial")

	return Validate.StructExcept(user)

}

func ValidateStruct(user interface{}) error {

	logger.Info("validating normal")
	return Validate.Struct(user)
}
