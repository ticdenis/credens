package bus

import (
	"credens/libs/shared/domain/bus"
	"credens/libs/shared/infrastructure"
	"errors"
)

type SyncCommandBus struct {
	commandHandlers []bus.CommandHandler
}

func NewSyncCommandBus(commandHandlers []interface{}) bus.CommandBus {
	var handlers []bus.CommandHandler

	for _, handler := range commandHandlers {
		if commandHandler, ok := handler.(bus.CommandHandler); ok {
			handlers = append(handlers, commandHandler)
		}
	}

	return &SyncCommandBus{handlers}
}

func (commandBus SyncCommandBus) Dispatch(command bus.Command) error {
	for _, handler := range commandBus.commandHandlers {
		if handler.SubscribedTo() == command.CommandName() {
			return handler.Execute(command)
		}
	}

	return infrastructure.NewInfrastructureError("404", "command not found", command, errors.New(""))
}
