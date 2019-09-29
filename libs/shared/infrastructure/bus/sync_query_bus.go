package bus

import (
	"credens/libs/shared/domain/bus"
	"credens/libs/shared/infrastructure"
	"errors"
)

type SyncQueryBus struct {
	queryHandlers []bus.QueryHandler
}

func NewSyncQueryBus(queryHandlers []bus.QueryHandler) *SyncQueryBus {
	return &SyncQueryBus{queryHandlers}
}

func (bus *SyncQueryBus) Ask(query bus.Query) (interface{}, error) {
	for _, handler := range bus.queryHandlers {
		if handler.SubscribedTo() == query.QueryName() {
			return handler.Execute(query)
		}
	}

	return nil, infrastructure.NewInfrastructureError("404", "query not found", query, errors.New(""))
}
