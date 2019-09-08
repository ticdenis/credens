package bus

import "credens/src/shared/domain/bus"

type SyncCommandBus struct {
	commandHandlers []bus.CommandHandler
}

func NewSyncCommandBus(commandHandlers []bus.CommandHandler) bus.CommandBus {
	return &SyncCommandBus{commandHandlers}
}

func (bus *SyncCommandBus) Dispatch(command bus.Command) {
	for _, handler := range bus.commandHandlers {
		if handler.SubscribedTo() == command.CommandName() {
			handler.Execute(command)
		}
	}
}
