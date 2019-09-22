package controller

import (
	"credens/src/application/create"
	"credens/src/domain/account"
	"credens/src/shared/application/serializer"
	"credens/src/shared/domain/bus"
	"credens/src/user_interface/http/contracts"
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func NewCreateAccountPostController(commandBus bus.CommandBus, jsonSerializer serializer.JSONSerializer) func(w http.ResponseWriter, r *http.Request) {
	validateCommandFromRequest := func(w http.ResponseWriter, r *http.Request) (bus.Command, *contracts.JSONAPIErrorObject) {
		var bodyParsed struct {
			Name     string `json:"name" validate:"required"`
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		if err := json.NewDecoder(r.Body).Decode(&bodyParsed); err != nil {
			return nil, contracts.NewJSONAPIErrorObject(err, http.StatusInternalServerError)
		}

		if err := validator.New().Struct(bodyParsed); err != nil {
			return nil, contracts.NewJSONAPIErrorObject(err, http.StatusBadRequest)
		}

		return *create.NewCreateAccountCommand(
			account.NewAccountId(nil).Value(),
			bodyParsed.Name,
			bodyParsed.Username,
			bodyParsed.Password,
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
