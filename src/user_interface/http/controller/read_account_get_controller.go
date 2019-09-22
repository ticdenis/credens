package controller

import (
	"credens/src/application/read"
	"credens/src/shared/application/serializer"
	"credens/src/shared/domain/bus"
	"credens/src/user_interface/http/contracts"
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
