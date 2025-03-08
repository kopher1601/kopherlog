package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"kopherlog/domain"
	"net/http"
	"regexp"
	"strconv"
)

var validate = validator.New()

func ValidateQueryParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams domain.PostSearchParams
		var paramErrors []*domain.ValidationError
		pageStr := c.Query("page")
		if !isNumeric(pageStr) {
			paramErrors = append(paramErrors, &domain.ValidationError{
				Field:   "page",
				Message: "page must be positive numbers",
			})
		}

		sizeStr := c.Query("size")
		if !isNumeric(pageStr) {
			paramErrors = append(paramErrors, &domain.ValidationError{
				Field:   "size",
				Message: "size must be positive numbers",
			})
		}

		if len(paramErrors) > 0 {
			resp := &domain.ErrorResponse{
				Code:        http.StatusBadRequest,
				Message:     "validation failed",
				Validations: paramErrors,
			}
			c.AbortWithStatusJSON(resp.Code, resp)
			return
		}

		queryParams.Page, _ = strconv.Atoi(pageStr)
		queryParams.Size, _ = strconv.Atoi(sizeStr)

		if err := validate.Struct(queryParams); err != nil {
			resp := &domain.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failed",
			}
			resp.AddValidationErrors(err)
			c.AbortWithStatusJSON(resp.Code, resp)
			return
		}

		c.Set("queryParams", queryParams)
	}
}

func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	numericRegex := regexp.MustCompile(`^\d+$`) // 数字のみの文字列か
	return numericRegex.MatchString(s)
}
