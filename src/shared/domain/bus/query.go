package bus

type Query interface {
	Message() Message
	QueryName() string
	Data() interface{}
}

var QueryMessageType = "query"

type QueryBus interface {
	Ask(query Query) (interface{}, error)
}

type QueryHandler interface {
	SubscribedTo() string
	Execute(query Query) (interface{}, error)
}
