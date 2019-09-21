package account

import (
	"credens/src/domain/account"
	coreError "credens/src/shared/domain/error"
)

type InMemoryAccountRepository struct {
	accounts []*account.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{[]*account.Account{}}
}

func (repo InMemoryAccountRepository) Add(account *account.Account) {
	repo.accounts = append(repo.accounts, account)
}

func (repo InMemoryAccountRepository) Search(id account.AccountId) (*account.Account, error) {
	for _, acc := range repo.accounts {
		if acc.Id().Value() == id.Value() {
			return acc, nil
		}
	}

	return nil, coreError.NewDomainError("404", "account not found", id)
}
