package read

import (
	"credens/src/domain/account"
	"credens/src/shared/domain/bus"
)

type ReadAccountQueryHandler struct {
	accountRepository account.AccountRepository
}

func NewReadAccountQueryHandler(
	accountRepository account.AccountRepository,
) *ReadAccountQueryHandler {
	return &ReadAccountQueryHandler{
		accountRepository,
	}
}

func (handler ReadAccountQueryHandler) SubscribedTo() string {
	return "read_account"
}

func (handler ReadAccountQueryHandler) Execute(query bus.Query) (interface{}, error) {
	id := account.NewAccountId(query.(ReadAccountQuery).data.Id)

	aggregate, err := handler.accountRepository.Search(id)
	if err != nil {
		return nil, err

	}

	return *NewReadAccountResponse(*aggregate), nil
}
