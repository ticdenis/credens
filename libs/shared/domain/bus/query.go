package bus

type (
	Query interface {
		Message() Message
		QueryName() string
		Data() interface{}
	}

	BaseQuery struct {
		queryName string
		message   Message
		data      interface{}
	}

	QueryBus interface {
		Ask(query Query) (interface{}, error)
	}

	QueryHandler interface {
		SubscribedTo() string
		Execute(query Query) (interface{}, error)
	}
)

func NewQuery(queryName string, data interface{}) *BaseQuery {
	return &BaseQuery{
		queryName: queryName,
		message:   *NewMessage("query"),
		data:      data,
	}
}

func (cmd BaseQuery) Message() Message {
	return cmd.message
}

func (cmd BaseQuery) QueryName() string {
	return cmd.queryName
}

func (cmd BaseQuery) Data() interface{} {
	return cmd.data
}
