package read

import (
	"credens/src/shared/domain/bus"
)

type ReadAccountQueryData struct {
	Id string
}

type ReadAccountQuery struct {
	message bus.Message
	data    ReadAccountQueryData
}

func NewReadAccountQuery(id string) *ReadAccountQuery {
	return &ReadAccountQuery{
		*bus.NewMessage(bus.QueryMessageType),
		ReadAccountQueryData{id},
	}
}

func (command ReadAccountQuery) Message() bus.Message {
	return command.message
}

func (command ReadAccountQuery) QueryName() string {
	return "read_account"
}

func (command ReadAccountQuery) Data() interface{} {
	return command.data
}
