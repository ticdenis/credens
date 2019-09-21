package account

type AccountBuilder struct {
}

func NewAccountBuilder() *AccountBuilder {
	return &AccountBuilder{}
}

func (builder *AccountBuilder) Build(id string, name string, username string, password string) *Account {
	return NewAccount(
		NewAccountId(id),
		NewAccountName(name),
		NewAccountUsername(username),
		NewAccountPassword(password),
	)
}
