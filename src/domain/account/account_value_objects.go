package account

import "credens/src/shared/domain/value_object"

type AccountId struct {
	value_object.UUID
}

func NewAccountId(value string) AccountId {
	return AccountId{UUID: *value_object.NewUuid(value)}
}

type AccountName struct {
	value_object.String
}

func NewAccountName(value string) AccountName {
	return AccountName{String: *value_object.NewString(value)}
}

type AccountUsername struct {
	value_object.String
}

func NewAccountUsername(value string) AccountUsername {
	return AccountUsername{String: *value_object.NewString(value)}
}

type AccountPassword struct {
	value_object.String
}

func NewAccountPassword(value string) AccountPassword {
	return AccountPassword{String: *value_object.NewString(value)}
}
