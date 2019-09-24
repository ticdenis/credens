package application

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

func (err ApplicationError) Code() string {
	return err.code
}

func (err ApplicationError) Msg() string {
	return err.message
}

func (err ApplicationError) Data() interface{} {
	return err.data
}

func (err ApplicationError) Err() error {
	return nil
}

func (err ApplicationError) Error() string {
	var format string
	var args []interface{}

	if err.Data() != nil {
		format = `%s ApplicationError: %s [data]: %v`
		args = []interface{}{err.Code(), err.Msg(), err.Data()}
	} else {
		format = `%s ApplicationError: %s`
		args = []interface{}{err.Code(), err.Msg()}
	}

	return errors.New(fmt.Sprintf(format, args...)).Error()
}
