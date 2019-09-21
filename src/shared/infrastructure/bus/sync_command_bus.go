package bus

import "credens/src/shared/domain/bus"

type SyncCommandBus struct {
	commandHandlers []bus.CommandHandler
}

func NewSyncCommandBus(commandHandlers []interface{}) bus.CommandBus {
	var handlers []bus.CommandHandler

	for _, handler := range commandHandlers {
		if _, ok := handler.(bus.CommandHandler); ok {
			handlers = append(handlers, handler.(bus.CommandHandler))
		}
	}

	return &SyncCommandBus{handlers}
}

func (bus *SyncCommandBus) Dispatch(command bus.Command) {
	for _, handler := range bus.commandHandlers {
		if handler.SubscribedTo() == command.CommandName {
			handler.Execute(command)
		}
	}
}
