package utility

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ErrorMessages(err interface{}) []string {
	errorMessages := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("Error on field %s, condition:%s", e.Field(), e.ActualTag())
		errorMessages = append(errorMessages, errorMessage)
	}
	return errorMessages
}
