package common

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func LoadValidator() {
	fmt.Println("Loading validator...")
	Validate = validator.New()
}

func GetValidationErrors(err error) []string {
	errors := err.(validator.ValidationErrors)
	var humanErrors []string
	var humanError string
	for _, err := range errors {
		fmt.Println(err.Tag())
		switch err.Tag() {
		case "required":
			humanError = fmt.Sprintf("%s is required",
				err.Field())
		case "email":
			humanError = fmt.Sprintf("%s is not valid email",
				err.Field())
		case "gte":
			humanError = fmt.Sprintf("%s value must be greater than %s",
				err.Field(), err.Param())
		case "lte":
			humanError = fmt.Sprintf("%s value must be lower than %s",
				err.Field(), err.Param())
		case "min":
			humanError = fmt.Sprintf("%s length must be greater than or equal to %s", err.Field(), err.Param())
		}
		humanErrors = append(humanErrors, humanError)
	}
	return humanErrors
}
