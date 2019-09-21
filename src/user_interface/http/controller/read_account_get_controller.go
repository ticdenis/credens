package controller

import (
	"credens/src/application/read"
	"credens/src/shared/domain/bus"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func NewReadAccountGetController(queryBus bus.QueryBus) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")

		if vars["id"] == "" {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, fmt.Sprintf(`{"data": null, "errors": ["%s"]}`, "id query param mandatory!"))

			return
		}

		res, err := queryBus.Ask(*read.NewReadAccountQuery(vars["id"]))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, fmt.Sprintf(`{"data": null, "errors": ["%s"]}`, err.Error()))

			return
		}

		resParsed, err := json.Marshal(&res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, fmt.Sprintf(`{"data": null, "errors": ["%s"]}`, err.Error()))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"data": "%s", "errors": []}`, resParsed)))
	}
}
