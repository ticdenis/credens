package controller

import (
	"credens/apps/http/contracts"
	"credens/libs/shared/application/serializer"
	"net/http"
)

func NewHealthzGetController(jsonSerializer serializer.JSONSerializer) (func(w http.ResponseWriter, r *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonResponder := contracts.NewJSONResponder(w, r, jsonSerializer)

		data := map[string]string{"status": "OK"}

		jsonResponder.DataResponse(http.StatusOK, "service", nil, data)
	}
}
