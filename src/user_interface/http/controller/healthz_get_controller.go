package controller

import (
	"credens/src/shared/application/serializer"
	"credens/src/user_interface/http/contracts"
	"net/http"
)

func NewHealthzGetController(jsonSerializer serializer.JSONSerializer) (func(w http.ResponseWriter, r *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonResponder := contracts.NewJSONResponder(w, r, jsonSerializer)

		data := map[string]string{"status": "OK"}

		jsonResponder.DataResponse(http.StatusOK, "service", nil, data)
	}
}
