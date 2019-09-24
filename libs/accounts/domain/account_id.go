package domain

import (
	value_object2 "credens/libs/shared/domain/value_object"
)

type AccountId struct {
	value_object2.UUID
}

func NewAccountId(value interface{}) AccountId {
	return AccountId{UUID: *value_object2.NewUuid(value)}
}
