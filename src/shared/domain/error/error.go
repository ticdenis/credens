package error

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

var domainErrorFormat = "DomainError [%s]: %s\n[data]: \"%v\""

func (err DomainError) Error() string {
	return errors.New(
		fmt.Sprintf(
			domainErrorFormat,
			err.code,
			err.message,
			err.data,
		),
	).Error()
}
