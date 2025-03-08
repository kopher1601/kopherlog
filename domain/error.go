package domain

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code        int                `json:"code"`
	Message     string             `json:"message"`
	Validations []*ValidationError `json:"validations"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var customMessages = map[string]string{
	"required": "必須です。",
	"lte":      "サイズが大きいすげです。",
}

func (e *ErrorResponse) AddValidationErrors(error error) {
	var validations []*ValidationError
	var validationErrors validator.ValidationErrors
	if errors.As(error, &validationErrors) {
		for _, e := range validationErrors {
			msg := e.Error()
			// カスタムメッセージがあれば交換
			if customMsg, exists := customMessages[e.Tag()]; exists {
				msg = customMsg
			}

			validations = append(validations, &ValidationError{
				Field:   e.Field(),
				Message: msg,
			})
		}
	}
	e.Validations = validations
}
