package create

import (
	"credens/src/domain/account"
	"credens/src/shared/domain/bus"
)

type CreateAccountCommandHandler struct {
	bus.CommandHandler
	accountRepository account.AccountRepository
	accountBuilder    account.AccountBuilder
	eventPublisher    bus.EventPublisher
}

func NewCreateAccountCommandHandler(
	accountRepository account.AccountRepository,
	accountBuilder account.AccountBuilder,
	eventPublisher bus.EventPublisher,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		*bus.NewCommandHandler(createAccountCommandName),
		accountRepository,
		accountBuilder,
		eventPublisher,
	}
}

func (handler *CreateAccountCommandHandler) Execute(command CreateAccountCommand) {
	aggregate := handler.accountBuilder.Build(
		command.Data.Id,
		command.Data.Name,
		command.Data.Username,
		command.Data.Password,
	)

	handler.accountRepository.Add(aggregate)

	handler.eventPublisher.Publish(aggregate.PullDomainEvents()...)
}
