package infrastructure

import (
	"fmt"
	"github.com/pkg/errors"
)

type InfrastructureError struct {
	code    string
	message string
	data    interface{}
	err     error
}

func NewInfrastructureError(code string, message string, data interface{}, err error) *InfrastructureError {
	return &InfrastructureError{code, message, data, err}
}

func (err InfrastructureError) Code() string {
	return err.code
}

func (err InfrastructureError) Msg() string {
	return err.message
}

func (err InfrastructureError) Data() interface{} {
	return err.data
}

func (err InfrastructureError) Err() error {
	return err.err
}

func (err InfrastructureError) Error() string {
	var format string
	var args []interface{}

	if err.Data() != nil {
		format = `%s InfrastructureError: %s [data]: %v`
		args = []interface{}{err.Code(), err.Msg(), err.Data()}
	} else {
		format = `%s InfrastructureError: %s`
		args = []interface{}{err.Code(), err.Msg()}
	}

	return errors.New(fmt.Sprintf(format, args...)).Error()
}

func PanicIfError(err error, msg string) {
	if err != nil {
		panic(errors.Wrap(err, msg))
	}
}
