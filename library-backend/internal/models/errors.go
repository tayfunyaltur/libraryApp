package models

import "net/http"

// Standard API Error Response
type ErrorResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	Code      string `json:"code,omitempty"`
	Details   string `json:"details,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// Validation Error Response
type ValidationErrorResponse struct {
	Success bool                    `json:"success"`
	Error   string                  `json:"error"`
	Fields  map[string]string       `json:"fields,omitempty"`
	Details []ValidationErrorDetail `json:"details,omitempty"`
}

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

// Standard Success Response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Pagination Response
type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

// Common HTTP Errors
var (
	ErrBookNotFound     = &ErrorResponse{Success: false, Error: "Book not found", Code: "BOOK_NOT_FOUND"}
	ErrInvalidBookID    = &ErrorResponse{Success: false, Error: "Invalid book ID", Code: "INVALID_BOOK_ID"}
	ErrInvalidRequest   = &ErrorResponse{Success: false, Error: "Invalid request format", Code: "INVALID_REQUEST"}
	ErrValidationFailed = &ErrorResponse{Success: false, Error: "Validation failed", Code: "VALIDATION_FAILED"}
	ErrInternalServer   = &ErrorResponse{Success: false, Error: "Internal server error", Code: "INTERNAL_ERROR"}
	ErrDatabaseError    = &ErrorResponse{Success: false, Error: "Database error", Code: "DATABASE_ERROR"}
)

// HTTP Status Code Mapping
func (e *ErrorResponse) StatusCode() int {
	switch e.Code {
	case "BOOK_NOT_FOUND":
		return http.StatusNotFound
	case "INVALID_BOOK_ID", "INVALID_REQUEST", "VALIDATION_FAILED":
		return http.StatusBadRequest
	case "DATABASE_ERROR", "INTERNAL_ERROR":
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
