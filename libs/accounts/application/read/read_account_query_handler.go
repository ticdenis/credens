package read

import (
	"credens/libs/accounts/domain"
	"credens/libs/shared/domain/bus"
)

type ReadAccountQueryHandler struct {
	accountRepository domain.AccountRepository
}

func NewReadAccountQueryHandler(
	accountRepository domain.AccountRepository,
) *ReadAccountQueryHandler {
	return &ReadAccountQueryHandler{accountRepository}
}

func (handler ReadAccountQueryHandler) SubscribedTo() string {
	return queryName
}

func (handler ReadAccountQueryHandler) Execute(query bus.Query) (interface{}, error) {
	data := query.Data().(ReadAccountQueryData)

	aggregate, err := handler.accountRepository.Search(domain.NewAccountId(data.Id))
	if err != nil {
		return nil, err
	}

	return *NewReadAccountResponse(*aggregate), nil
}
