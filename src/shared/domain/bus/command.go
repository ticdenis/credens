package bus

import "credens/src/shared/domain/value_object"

type Command struct {
	messageId   string
	messageType string
	commandName string
}

func NewCommand(commandName string) *Command {
	return &Command{
		value_object.UUID(nil).Value(),
		"command",
		commandName,
	}
}

func (command *Command) MessageId() string {
	return command.messageId
}

func (command *Command) MessageType() string {
	return command.messageType
}

func (command *Command) CommandName() string {
	return command.commandName
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