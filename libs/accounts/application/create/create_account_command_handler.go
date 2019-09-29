package create

import (
	"credens/libs/accounts/domain"
	"credens/libs/shared/domain/bus"
)

type CreateAccountCommandHandler struct {
	accountRepository domain.AccountRepository
	eventPublisher    bus.EventPublisher
}

func NewCreateAccountCommandHandler(
	accountRepository domain.AccountRepository,
	eventPublisher bus.EventPublisher,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		accountRepository,
		eventPublisher,
	}
}

func (handler CreateAccountCommandHandler) SubscribedTo() string {
	return commandName
}

func (handler CreateAccountCommandHandler) Execute(command bus.Command) error {
	data := command.Data().(CreateAccountCommandData)

	aggregate := domain.NewAccount(
		domain.NewAccountId(data.Id),
		domain.NewAccountName(data.Name),
		domain.NewAccountUsername(data.Username),
		domain.NewAccountPassword(data.Password),
	)

	if err := handler.accountRepository.Add(aggregate); err != nil {
		return err
	}

	handler.eventPublisher.Publish(aggregate.PullDomainEvents()...)

	return nil
}
