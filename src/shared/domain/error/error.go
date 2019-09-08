package error

import "errors"

type Error struct {
	code    string
	message string
}

func NewDomainError(code string, message string) *Error {
	return &Error{code, message}
}

// Implements errors.Error
func (domainError *Error) Error() string {
	return errors.New(domainError.code + " -> " + domainError.message).Error()
}
