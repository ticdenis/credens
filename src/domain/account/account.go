package account

import (
	"credens/src/shared/domain/aggregate"
)

type Account struct {
	aggregate.AggregateRoot
	id       AccountId
	name     AccountName
	username AccountUsername
	password AccountPassword
}

func NewAccount(id AccountId, name AccountName, username AccountUsername, password AccountPassword) *Account {
	account := &Account{id: id, name: name, username: username, password: password}

	account.RecordDomainEvent(
		*NewAccountCreated(id.Value(), name.Value(), username.Value()),
	)

	return account
}

func (account Account) Id() AccountId {
	return account.id
}

func (account Account) Name() AccountName {
	return account.name
}

func (account Account) Username() AccountUsername {
	return account.username
}

func (account Account) Password() AccountPassword {
	return account.password
}
