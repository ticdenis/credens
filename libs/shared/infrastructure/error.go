package infrastructure

import (
	"errors"
	"fmt"
)

type InfrastructureError struct {
	Code    string
	Message string
	Data    interface{}
	Err     error
}

func NewInfrastructureError(code string, message string, data interface{}, err error) *InfrastructureError {
	return &InfrastructureError{code, message, data, err}
}

var infrastructureErrorFormat = "InfrastructureError [%s]: %s\n[data]: \"%v\"\n[err]: %s"

func (err InfrastructureError) Error() string {
	return errors.New(
		fmt.Sprintf(
			infrastructureErrorFormat,
			err.Code,
			err.Message,
			err.Data,
			err.Err.Error(),
		),
	).Error()
}
