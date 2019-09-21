package bus

import "credens/src/shared/domain/bus"

type SyncQueryBus struct {
	queryHandlers []bus.QueryHandler
}

func NewSyncQueryBus(queryHandlers []interface{}) bus.QueryBus {
	var handlers []bus.QueryHandler

	for _, handler := range queryHandlers {
		if _, ok := handler.(bus.QueryHandler); ok {
			handlers = append(handlers, handler.(bus.QueryHandler))
		}
	}

	return &SyncQueryBus{handlers}
}

func (bus *SyncQueryBus) Ask(query bus.Query) bus.Response {
	for _, handler := range bus.queryHandlers {
		if handler.SubscribedTo() == query.QueryName() {
			return handler.Execute(query)
		}
	}
	return nil
}
