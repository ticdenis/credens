package handler

import (
	"credens/libs/accounts/application/read"
	"credens/libs/shared/domain/bus"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReadAccountGetHandler struct {
	queryBus bus.QueryBus
	query    bus.Query
}

func NewReadAccountGetHandler(queryBus bus.QueryBus) *ReadAccountGetHandler {
	return &ReadAccountGetHandler{queryBus: queryBus}
}

func (handler *ReadAccountGetHandler) validate(context *gin.Context) *ResponseError {
	accountId := context.Param("id")

	if accountId == "" {
		return NewResponseError(errors.New("URL param 'id' is mandatory"))
	}

	handler.query = *read.NewReadAccountQuery(read.ReadAccountQueryData{Id: accountId})

	return nil
}

func (handler ReadAccountGetHandler) Handle(context *gin.Context) (*Response, *ResponseError) {
	if err := handler.validate(context); err != nil {
		return nil, err
	}

	res, err := handler.queryBus.Ask(handler.query)
	if err != nil {
		return nil, NewResponseError(err)
	}

	return NewResponse(
		http.StatusOK,
		map[string]interface{}{
			"data": map[string]interface{}{
				"type":       "accounts",
				"id":         handler.query.Data().(read.ReadAccountQueryData).Id,
				"attributes": res,
			},
		},
	), nil
}
