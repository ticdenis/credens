package account

import "credens/src/domain/account"

type InMemoryAccountRepository struct {
	accounts []account.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{[]account.Account{}}
}

func (repo InMemoryAccountRepository) Add(account account.Account) {
	repo.accounts = append(repo.accounts, account)
}
