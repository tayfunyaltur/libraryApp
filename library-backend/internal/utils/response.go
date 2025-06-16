package utils

import (
	"library-backend/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// SendError sends a standardized error response
func SendError(c *gin.Context, statusCode int, message, code string, details ...string) {
	errorResp := &models.ErrorResponse{
		Success:   false,
		Error:     message,
		Code:      code,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if len(details) > 0 {
		errorResp.Details = details[0]
	}

	c.JSON(statusCode, errorResp)
}

// SendSuccess sends a standardized success response
func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	response := &models.SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	c.JSON(statusCode, response)
}

// SendValidationError sends validation error response
func SendValidationError(c *gin.Context, err error) {
	validationErrors := make(map[string]string)

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, validationErr := range validationErrs {
			validationErrors[validationErr.Field()] = getValidationMessage(validationErr)
		}
	}

	response := &models.ValidationErrorResponse{
		Success: false,
		Error:   "Validation failed",
		Fields:  validationErrors,
	}

	c.JSON(http.StatusBadRequest, response)
}

func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "url":
		return "Invalid URL format"
	case "oneof":
		return "Invalid value"
	default:
		return "Invalid value"
	}
}
