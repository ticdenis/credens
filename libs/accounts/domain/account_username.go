package domain

import (
	value_object2 "credens/libs/shared/domain/value_object"
)

type AccountUsername struct {
	value_object2.String
}

func NewAccountUsername(value string) AccountUsername {
	return AccountUsername{String: *value_object2.NewString(value)}
}
