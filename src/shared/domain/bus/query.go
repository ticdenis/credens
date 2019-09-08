package bus

type Query struct {
	queryName string
}

func NewQuery(queryName string) *Query {
	return &Query{queryName}
}

func (query *Query) QueryName() string {
	return query.queryName
}

type QueryBus interface {
	Ask(query Query) Response
}

type QueryHandler interface {
	SubscribedTo() string
	Execute(query Query) Response
}

type Response interface {
}
