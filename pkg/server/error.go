package server

import "net/http"

var (
	ErrorGenericError     = NewError("Something expected went wrong.", "GENERIC_ERROR", http.StatusInternalServerError)
	ErrorInvalidAuthToken = NewError("Your request has an invalid authorization token.", "AUTH_TOKEN_INVALID", http.StatusUnauthorized)
	ErrorMissingAuthToken = NewError("Your request is missing an authorization token.", "AUTH_TOKEN_MISSING", http.StatusUnauthorized)
	ErrorNotFound         = NewError("The entity you requested was not found.", "NOT_FOUND", http.StatusNotFound)
	ErrorValidationErrors = NewError("Your request has some validation errors.", "VALIDATION_ERRORS", http.StatusBadRequest)
)

type ErrorInterface interface {
	GetMessage() string
	GetStatus() int
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Type    string `json:"type"`
}

func NewError(message, errorType string, status int) *Error {
	return &Error{
		Message: message,
		Status:  status,
		Type:    errorType,
	}
}

func (e Error) GetMessage() string {
	return e.Message
}

func (e Error) GetStatus() int {
	return e.Status
}

func (e Error) String() string {
	return e.Message
}

type ValidationErrors struct {
	Fields []ValidationFieldError `json:"fields"`
	*Error
}

type ValidationFieldError struct {
	Field   string                 `json:"field"`
	Fields  []ValidationFieldError `json:"fields,omitempty"`
	Message string                 `json:"message"`
}

func NewValidationErrors(fields []ValidationFieldError) *ValidationErrors {
	return &ValidationErrors{
		Error:  ErrorValidationErrors,
		Fields: fields,
	}
}
