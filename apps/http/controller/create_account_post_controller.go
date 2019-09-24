package controller

import (
	"credens/apps/http/contracts"
	"credens/libs/accounts/application/create"
	"credens/libs/accounts/domain"
	"credens/libs/shared/application/serializer"
	"credens/libs/shared/domain/bus"
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func NewCreateAccountPostController(commandBus bus.CommandBus, jsonSerializer serializer.JSONSerializer) func(w http.ResponseWriter, r *http.Request) {
	validateCommandFromRequest := func(w http.ResponseWriter, r *http.Request) (bus.Command, *contracts.JSONAPIErrorObject) {
		var body struct {
			Name     string `json:"name" validate:"required"`
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return nil, contracts.NewJSONAPIErrorObject(err, http.StatusInternalServerError)
		}

		if err := validator.New().Struct(body); err != nil {
			return nil, contracts.NewJSONAPIErrorObject(err, http.StatusBadRequest)
		}

		return *create.NewCreateAccountCommand(
			domain.NewAccountId(nil).Value(),
			body.Name,
			body.Username,
			body.Password,
		), nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		jsonResponder := contracts.NewJSONResponder(w, r, jsonSerializer)

		command, err := validateCommandFromRequest(w, r)
		if err != nil {
			jsonResponder.ErrorsResponse(err.HttpStatus, *err)
			return
		}

		if err := commandBus.Dispatch(command); err != nil {
			jsonResponder.ErrorsResponse(http.StatusConflict, err)
			return
		}

		jsonResponder.DataResponse(
			http.StatusCreated,
			"accounts",
			command.Data().(create.CreateAccountCommandData).Id,
			nil,
		)
	}
}
