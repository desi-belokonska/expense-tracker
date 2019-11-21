package main

// ErrorJSON is a stuct which describes the errors returned to the client
type ErrorJSON struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Errors  []ErrorField `json:"errors,omitempty"`
}

// ErrorField gives specific information about what was wrong in each field
type ErrorField struct {
	Message string
	Field   string
}

// InvalidInput is the code used in the error for all types of invalid input
const InvalidInput = "INVALID_INPUT"

// NotFound is the code for when the resource doesn't exist
const NotFound = "NOT_FOUND"

// CreateErrorInvalidInput returns an initialised error with code INVALID_INPUT
// and the given message and errors
func CreateErrorInvalidInput(message string, errors []ErrorField) ErrorJSON {
	return ErrorJSON{InvalidInput, message, errors}
}

// CreateErrorNotFound returns an initialised error with code NOT_FOUND
// and the given message
func CreateErrorNotFound(message string) ErrorJSON {
	return ErrorJSON{NotFound, message, nil}
}
