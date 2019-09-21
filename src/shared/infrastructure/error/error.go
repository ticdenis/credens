package error

import (
	"errors"
	"fmt"
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

var infrastructureErrorFormat = "InfrastructureError [%s]: %s\n[data]: \"%v\"\n[err]: %s"

func (err InfrastructureError) Error() string {
	return errors.New(
		fmt.Sprintf(
			infrastructureErrorFormat,
			err.code,
			err.message,
			err.data,
			err.err.Error(),
		),
	).Error()
}
