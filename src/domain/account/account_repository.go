package account

type AccountRepository interface {
	Add(account *Account)
	Search(id AccountId) (*Account, error)
}
