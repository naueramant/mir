package config

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

var (
	syntaxes = []string{"v1"}
)

func Validate(t Configuration) error {
	validate := validator.New()

	validate.RegisterValidation("syntax", validateSyntax)

	err := validate.Struct(t)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Syntax":
				return errors.New("Malformed config, " + err.StructNamespace() + ": " + fmt.Sprintf("%v", err.Value()) + " is not a valid syntax version")
			default:
				return errors.New("Malformed config, " + err.StructNamespace() + ": " + strings.ToLower(fmt.Sprintf("%v", err.StructField())) + " is required")
			}
		}
	}

	return err
}

func validateSyntax(fl validator.FieldLevel) bool {
	return contains(syntaxes, fl.Field().String())
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
