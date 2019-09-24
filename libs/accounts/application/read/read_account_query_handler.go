package read

import (
	"credens/libs/accounts/domain"
	bus2 "credens/libs/shared/domain/bus"
)

type ReadAccountQueryHandler struct {
	svc ReadAccountService
}

func NewReadAccountQueryHandler(
	readAccountService ReadAccountService,
) *ReadAccountQueryHandler {
	return &ReadAccountQueryHandler{readAccountService}
}

func (handler ReadAccountQueryHandler) SubscribedTo() string {
	return queryName
}

func (handler ReadAccountQueryHandler) Execute(query bus2.Query) (interface{}, error) {
	data := query.Data().(ReadAccountQueryData)

	return handler.svc.Execute(
		domain.NewAccountId(data.Id),
	)
}
