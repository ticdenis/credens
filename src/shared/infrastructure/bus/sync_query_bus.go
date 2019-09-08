package bus

import "credens/src/shared/domain/bus"

type SyncQueryBus struct {
	queryHandlers []bus.QueryHandler
}

func NewSyncQueryBus(queryHandlers []bus.QueryHandler) bus.QueryBus {
	return &SyncQueryBus{queryHandlers}
}

func (bus *SyncQueryBus) Ask(query bus.Query) bus.Response {
	for _, handler := range bus.queryHandlers {
		if handler.SubscribedTo() == query.QueryName() {
			return handler.Execute(query)
		}
	}
	return nil
}
