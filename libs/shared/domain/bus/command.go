package bus

type Command interface {
	Message() Message
	CommandName() string
	Data() interface{}
}

var CommandMessageType = "command"

type CommandBus interface {
	Dispatch(command Command) error
}

type CommandHandler interface {
	SubscribedTo() string
	Execute(command Command) error
}
