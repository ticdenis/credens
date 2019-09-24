package read

import (
	"credens/libs/accounts/domain"
)

type ReadAccountService struct {
	accountRepository domain.AccountRepository
}

func NewReadAccountService(accountRepository domain.AccountRepository) *ReadAccountService {
	return &ReadAccountService{accountRepository}
}

func (handler ReadAccountService) SubscribedTo() string {
	return queryName
}

func (handler ReadAccountService) Execute(id domain.AccountId) (interface{}, error) {
	aggregate, err := handler.accountRepository.Search(id)
	if err != nil {
		return nil, err
	}

	return *NewReadAccountResponse(*aggregate), nil
}
