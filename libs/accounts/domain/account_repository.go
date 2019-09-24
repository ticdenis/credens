package domain

type AccountRepository interface {
	Add(account *Account) error
	Search(id AccountId) (*Account, error)
}
