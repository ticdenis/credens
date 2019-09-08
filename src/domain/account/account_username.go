package account

import "credens/src/shared/domain/value_object"

type AccountUsername struct {
	value_object.String
}

func NewAccountUsername(value string) AccountUsername {
	return AccountUsername{String: *value_object.NewString(value)}
}
