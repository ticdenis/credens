package create

import (
	"credens/libs/accounts/domain"
	bus2 "credens/libs/shared/domain/bus"
)

type CreateAccountService struct {
	accountRepository domain.AccountRepository
	eventPublisher    bus2.EventPublisher
}

func NewCreateAccountService(
	accountRepository domain.AccountRepository,
	eventPublisher bus2.EventPublisher,
) *CreateAccountService {
	return &CreateAccountService{
		accountRepository,
		eventPublisher,
	}
}

func (svc CreateAccountService) Execute(
	id domain.AccountId,
	name domain.AccountName,
	username domain.AccountUsername,
	password domain.AccountPassword,
) (err error) {
	aggregate := domain.NewAccount(id, name, username, password)

	if err = svc.accountRepository.Add(aggregate); err != nil {
		return err
	}

	svc.eventPublisher.Publish(aggregate.PullDomainEvents()...)

	return nil
}
