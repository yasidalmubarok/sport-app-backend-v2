package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errorsList []string

	// Check if the error is of type ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errorsList = append(errorsList, e.Error())
		}
		return errorsList
	}

	// Check if the error is a JSON UnmarshalTypeError
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		errorsList = append(errorsList, fmt.Sprintf("Field '%s' has an invalid type", unmarshalTypeError.Field))
		return errorsList
	}

	// Check if the error is a JSON SyntaxError
	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		errorsList = append(errorsList, "Invalid JSON syntax")
		return errorsList
	}

	// Handle general errors
	errorsList = append(errorsList, err.Error())
	return errorsList
}

func ValidatePhoneNumber(phone string) Error {
	re := regexp.MustCompile(`^(08|628)[0-9]{8,11}$`)
	if !re.MatchString(phone) {
		return NewBadRequestError("invalid phone number")
	}
	return nil
}

func ValidateEmail(email string) Error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return NewBadRequestError("invalid email")
	}
	return nil
}

func ValidateUsername(username string) Error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{8,18}$`)
	if !re.MatchString(username) {
		return NewBadRequestError("invalid username")
	}
	return nil
}