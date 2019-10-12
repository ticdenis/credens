package domain

import "credens/libs/shared/domain/value_object"

type AccountId struct {
	value_object.UUID
}

func NewAccountId(value interface{}) AccountId {
	return AccountId{UUID: *value_object.NewUuid(value)}
}
