package controller

import (
	"io"
	"net/http"
)

func NewHealthzGetController() (func(w http.ResponseWriter, r *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, `{"status": "OK"}`)
	}

}
