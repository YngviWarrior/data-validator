package validator_test

import (
	"testing"

	validator "github.com/YngviWarrior/data-validator"
)

func Test(t *testing.T) {
	validationHandler := validator.NewValidator()

	var obj struct {
		Date string `validate:"date"`
	}

	obj.Date = "165asd"

	err := validationHandler.Validation(obj)
	if err != nil {
		t.Error(err)
	}
}
