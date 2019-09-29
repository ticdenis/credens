package mocks

import (
	"credens/libs/accounts/domain"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (mock *AccountRepositoryMock) Add(account *domain.Account) error {
	args := mock.Called(account)

	return args.Error(0)
}

func (mock *AccountRepositoryMock) Search(id domain.AccountId) (*domain.Account, error) {
	args := mock.Called(id)

	return args.Get(0).(*domain.Account), args.Error(1)
}
