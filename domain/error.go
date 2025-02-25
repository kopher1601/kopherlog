package domain

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var customMessages = map[string]string{
	"required": "必須です。",
}

func GenerateErrorResponse(error error) []ErrorResponse {
	var errors []ErrorResponse
	if validationErrors, ok := error.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			var msg string

			// カスタムメッセージがあれば交換
			if customMsg, exists := customMessages[e.Tag()]; exists {
				msg = customMsg
			}

			errors = append(errors, ErrorResponse{
				Field:   e.Field(),
				Message: msg,
			})
		}
	}
	return errors
}
