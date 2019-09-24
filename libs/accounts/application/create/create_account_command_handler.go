package create

import (
	"credens/libs/accounts/domain"
	bus2 "credens/libs/shared/domain/bus"
)

type CreateAccountCommandHandler struct {
	svc CreateAccountService
}

func NewCreateAccountCommandHandler(createAccountService CreateAccountService) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{createAccountService}
}

func (handler CreateAccountCommandHandler) SubscribedTo() string {
	return commandName
}

func (handler CreateAccountCommandHandler) Execute(command bus2.Command) error {
	data := command.Data().(CreateAccountCommandData)

	return handler.svc.Execute(
		domain.NewAccountId(data.Id),
		domain.NewAccountName(data.Name),
		domain.NewAccountUsername(data.Username),
		domain.NewAccountPassword(data.Password),
	)
}
