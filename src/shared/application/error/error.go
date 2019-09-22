package error

import (
	"errors"
	"fmt"
)

type ApplicationError struct {
	Code    string
	Message string
	Data    interface{}
}

func NewApplicationError(code string, message string, data interface{}) *ApplicationError {
	return &ApplicationError{code, message, data}
}

var applicationErrorFormat = "ApplicationError [%s]: %s\n[data]: \"%v\""

func (err ApplicationError) Error() string {
	return errors.New(
		fmt.Sprintf(
			applicationErrorFormat,
			err.Code,
			err.Message,
			err.Data,
		),
	).Error()
}
