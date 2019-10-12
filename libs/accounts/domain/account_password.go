package domain

import (
	"credens/libs/shared/domain/value_object"
)

type AccountPassword struct {
	value_object.String
}

func NewAccountPassword(value string) AccountPassword {
	return AccountPassword{String: *value_object.NewString(value)}
}
