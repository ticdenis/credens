package create

import (
	"credens/src/domain/account"
	"credens/src/shared/domain/bus"
)

type CreateAccountCommandHandler struct {
	bus.CommandHandler
	accountRepository account.AccountRepository
	eventPublisher    bus.EventPublisher
}

func NewCreateAccountCommandHandler(
	accountRepository account.AccountRepository,
	eventPublisher bus.EventPublisher,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		*bus.NewCommandHandler("create_account"),
		accountRepository,
		eventPublisher,
	}
}

func (handler *CreateAccountCommandHandler) Execute(command CreateAccountCommand) {
	account := account.NewAccount(
		account.NewAccountId(command.Data.Id),
		account.NewAccountName(command.Data.Name),
		account.NewAccountUsername(command.Data.UserName),
		account.NewAccountPassword(command.Data.Password),
	)

	handler.accountRepository.Add(*account)

	handler.eventPublisher.Publish(account.PullDomainEvents()...)
}
