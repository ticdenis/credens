package error

import (
	"errors"
	"fmt"
)

type ApplicationError struct {
	code    string
	message string
	data    interface{}
}

func NewApplicationError(code string, message string, data interface{}) *ApplicationError {
	return &ApplicationError{code, message, data}
}

var applicationErrorFormat = "ApplicationError [%s]: %s\n[data]: \"%v\""

func (err ApplicationError) Error() string {
	return errors.New(
		fmt.Sprintf(
			applicationErrorFormat,
			err.code,
			err.message,
			err.data,
		),
	).Error()
}
