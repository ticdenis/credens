package read

import (
	"credens/libs/shared/domain/bus"
)

var queryName = "read_account"

type (
	ReadAccountQuery struct {
		bus.BaseQuery
	}

	ReadAccountQueryData struct {
		Id string
	}
)

func NewReadAccountQuery(data ReadAccountQueryData) *ReadAccountQuery {
	return &ReadAccountQuery{*bus.NewQuery(queryName, data)}
}
