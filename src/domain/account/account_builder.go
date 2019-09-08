package account

type AccountBuilder struct {
}

func (builder *AccountBuilder) Build(id string, name string, username string, password string) *Account {
	return NewAccount(
		NewAccountId(id),
		NewAccountName(name),
		NewAccountUsername(username),
		NewAccountPassword(password),
	)
}
