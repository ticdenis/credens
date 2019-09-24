package domain

import (
	value_object2 "credens/libs/shared/domain/value_object"
)

type AccountPassword struct {
	value_object2.String
}

func NewAccountPassword(value string) AccountPassword {
	return AccountPassword{String: *value_object2.NewString(value)}
}
