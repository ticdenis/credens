package controller

import (
	"credens/src/application/create"
	"credens/src/domain/account"
	"credens/src/shared/domain/bus"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
)

func NewCreateAccountPostController(commandBus bus.CommandBus) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyParsed struct {
			Name     string `json:"name" validate:"required"`
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewDecoder(r.Body).Decode(&bodyParsed); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, fmt.Sprintf(`{"data": null, "errors": ["%s"]}`, err.Error()))

			return
		}

		if err := validator.New().Struct(bodyParsed); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, fmt.Sprintf(`{"data": null, "errors": ["%s"]}`, err.Error()))

			return
		}

		err := commandBus.Dispatch(*create.NewCreateAccountCommand(
			account.NewAccountId(nil).Value(),
			bodyParsed.Name,
			bodyParsed.Username,
			bodyParsed.Password,
		))
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			io.WriteString(w, fmt.Sprintf(`{"data": null, "errors": ["%s"]}`, err.Error()))

			return
		}

		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "{}")
	}
}
