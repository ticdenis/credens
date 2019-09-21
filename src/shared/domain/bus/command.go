package bus

type Command struct {
	Message
	CommandName string
}

var commandMessageType = "command"

func NewCommand(commandName string) *Command {
	return &Command{
		*NewMessage(commandMessageType),
		commandName,
	}
}

type CommandBus interface {
	Dispatch(command Command)
}

type CommandHandler struct {
	commandName string
}

func NewCommandHandler(commandName string) *CommandHandler {
	return &CommandHandler{commandName}
}

func (handler *CommandHandler) SubscribedTo() string {
	return handler.commandName
}

func (handler *CommandHandler) Execute(command Command) {
}
