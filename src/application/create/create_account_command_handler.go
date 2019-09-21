package create

import (
	"credens/src/domain/account"
	"credens/src/shared/domain/bus"
)

type CreateAccountCommandHandler struct {
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
		accountRepository,
		accountBuilder,
		eventPublisher,
	}
}

func (handler CreateAccountCommandHandler) SubscribedTo() string {
	return "create_account"
}

func (handler CreateAccountCommandHandler) Execute(command bus.Command) error {
	if cmd, ok := command.(CreateAccountCommand); ok {
		return handler.execute(cmd)
	}

	return nil
}

func (handler CreateAccountCommandHandler) execute(command CreateAccountCommand) error {
	aggregate := handler.accountBuilder.Build(
		command.data.Id,
		command.data.Name,
		command.data.Username,
		command.data.Password,
	)

	handler.accountRepository.Add(aggregate)

	handler.eventPublisher.Publish(aggregate.PullDomainEvents()...)

	return nil
}
