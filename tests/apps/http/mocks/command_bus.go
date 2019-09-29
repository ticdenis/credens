package mocks

import (
	"credens/libs/shared/domain/bus"
	"errors"
)

type TestCommandBus struct {
	commandHandlers []bus.CommandHandler
}

func NewTestCommandBus(commandHandlers []bus.CommandHandler) *TestCommandBus {
	return &TestCommandBus{commandHandlers}
}

func (commandBus TestCommandBus) Dispatch(command bus.Command) error {
	for _, handler := range commandBus.commandHandlers {
		if handler.SubscribedTo() == command.CommandName() {
			return handler.Execute(command)
		}
	}

	return errors.New("CommandHandler to execute Command received not found!")
}

