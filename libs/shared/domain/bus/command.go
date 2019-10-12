package bus

type (
	Command interface {
		Message() Message
		CommandName() string
		Data() interface{}
	}

	BaseCommand struct {
		commandName string
		message     Message
		data        interface{}
	}

	CommandBus interface {
		Dispatch(command Command) error
	}

	CommandHandler interface {
		SubscribedTo() string
		Execute(command Command) error
	}
)

func NewCommand(commandName string, data interface{}) *BaseCommand {
	return &BaseCommand{
		commandName: commandName,
		message:     *NewMessage("command"),
		data:        data,
	}
}

func (cmd BaseCommand) Message() Message {
	return cmd.message
}

func (cmd BaseCommand) CommandName() string {
	return cmd.commandName
}

func (cmd BaseCommand) Data() interface{} {
	return cmd.data
}
