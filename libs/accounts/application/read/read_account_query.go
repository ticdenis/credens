package read

import (
	bus2 "credens/libs/shared/domain/bus"
)

var queryName = "read_account"

type ReadAccountQueryData struct {
	Id string
}

type ReadAccountQuery struct {
	message bus2.Message
	data    ReadAccountQueryData
}

func NewReadAccountQuery(id string) *ReadAccountQuery {
	return &ReadAccountQuery{
		*bus2.NewMessage(bus2.QueryMessageType),
		ReadAccountQueryData{id},
	}
}

func (query ReadAccountQuery) Message() bus2.Message {
	return query.message
}

func (query ReadAccountQuery) QueryName() string {
	return queryName
}

func (query ReadAccountQuery) Data() interface{} {
	return query.data
}
