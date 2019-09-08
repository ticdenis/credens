package account

import "credens/src/shared/domain/value_object"

type AccountId struct {
	value_object.UUID
}

func NewAccountId(value string) AccountId {
	return AccountId{UUID: *value_object.NewUuid(value)}
}
