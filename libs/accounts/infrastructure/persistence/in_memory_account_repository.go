package persistence

import (
	domain2 "credens/libs/accounts/domain"
	"credens/libs/shared/domain"
)

type InMemoryAccountRepository struct {
	accounts []*domain2.Account
}

func NewInMemoryAccountRepository(accounts []*domain2.Account) *InMemoryAccountRepository {
	return &InMemoryAccountRepository{accounts}
}

func (repo InMemoryAccountRepository) Add(account *domain2.Account) error {
	repo.accounts = append(repo.accounts, account)

	return nil
}

func (repo InMemoryAccountRepository) Search(id domain2.AccountId) (*domain2.Account, error) {
	for _, acc := range repo.accounts {
		if acc.Id().Value() == id.Value() {
			return acc, nil
		}
	}

	return nil, domain.NewDomainError(
		"404",
		"accounts not found",
		map[string]interface{}{
			"id": id.Value(),
		},
	)
}
