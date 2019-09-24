package controller

import (
	"credens/apps/http/contracts"
	"credens/libs/accounts/application/read"
	"credens/libs/shared/application/serializer"
	"credens/libs/shared/domain/bus"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func NewReadAccountGetController(queryBus bus.QueryBus, jsonSerializer serializer.JSONSerializer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonResponder := contracts.NewJSONResponder(w, r, jsonSerializer)
		accountId := mux.Vars(r)["id"]

		// validation
		if accountId == "" {
			jsonResponder.ErrorsResponse(http.StatusBadRequest, errors.New("query param 'id' mandatory"))
			return
		}

		// action
		res, err := queryBus.Ask(*read.NewReadAccountQuery(accountId))
		if err != nil {
			jsonResponder.ErrorsResponse(http.StatusNotFound, err)
			return
		}

		jsonResponder.DataResponse(http.StatusOK, "accounts", accountId, res)
	}
}
