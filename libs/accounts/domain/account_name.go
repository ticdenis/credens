package domain

import (
	value_object2 "credens/libs/shared/domain/value_object"
)

type AccountName struct {
	value_object2.String
}

func NewAccountName(value string) AccountName {
	return AccountName{String: *value_object2.NewString(value)}
}
