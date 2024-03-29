package domain

import (
	"credens/libs/shared/domain/value_object"
)

type AccountName struct {
	value_object.String
}

func NewAccountName(value string) AccountName {
	return AccountName{String: *value_object.NewString(value)}
}
