package bus

import (
	"credens/src/shared/domain/bus"
	coreError "credens/src/shared/infrastructure/error"
	"errors"
)

type SyncQueryBus struct {
	queryHandlers []bus.QueryHandler
}

func NewSyncQueryBus(queryHandlers []interface{}) bus.QueryBus {
	var handlers []bus.QueryHandler

	for _, handler := range queryHandlers {
		if queryHandler, ok := handler.(bus.QueryHandler); ok {
			handlers = append(handlers, queryHandler)
		}
	}

	return &SyncQueryBus{handlers}
}

func (bus *SyncQueryBus) Ask(query bus.Query) (interface{}, error) {
	for _, handler := range bus.queryHandlers {
		if handler.SubscribedTo() == query.QueryName() {
			return handler.Execute(query)
		}
	}

	return nil, coreError.NewInfrastructureError("404", "query not found", query, errors.New(""))
}
