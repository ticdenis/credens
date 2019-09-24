package domain

import (
	"errors"
	"fmt"
)

type DomainError struct {
	code    string
	message string
	data    interface{}
}

func NewDomainError(code string, message string, data interface{}) *DomainError {
	return &DomainError{code, message, data}
}

func (err DomainError) Code() string {
	return err.code
}

func (err DomainError) Msg() string {
	return err.message
}

func (err DomainError) Data() interface{} {
	return err.data
}

func (err DomainError) Err() error {
	return nil
}

func (err DomainError) Error() string {
	var format string
	var args []interface{}

	if err.Data() != nil {
		format = `%s DomainError: %s [data]: %v`
		args = []interface{}{err.Code(), err.Msg(), err.Data()}
	} else {
		format = `%s DomainError: %s`
		args = []interface{}{err.Code(), err.Msg()}
	}

	return errors.New(fmt.Sprintf(format, args...)).Error()
}
