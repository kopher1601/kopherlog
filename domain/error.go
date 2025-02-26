package domain

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code        int               `json:"code"`
	Message     string            `json:"message"`
	Validations []ValidationError `json:"validations"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var customMessages = map[string]string{
	"required": "必須です。",
}

func (e *ErrorResponse) AddValidationErrors(error error) {
	var validations []ValidationError
	if validationErrors, ok := error.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			var msg string

			// カスタムメッセージがあれば交換
			if customMsg, exists := customMessages[e.Tag()]; exists {
				msg = customMsg
			}

			validations = append(validations, ValidationError{
				Field:   e.Field(),
				Message: msg,
			})
		}
	}
	e.Validations = validations
}
