package handler

import (
	"credens/libs/accounts/application/create"
	"credens/libs/accounts/domain"
	"credens/libs/shared/domain/bus"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateAccountPostHandler struct {
	commandBus bus.CommandBus
	command    bus.Command
}

func NewCreateAccountPostHandler(commandBus bus.CommandBus) *CreateAccountPostHandler {
	return &CreateAccountPostHandler{commandBus: commandBus}
}

func (handler *CreateAccountPostHandler) validate(context *gin.Context) *ResponseError {
	var json = struct {
		Id       string
		Name     string `json:"name" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{Id: domain.NewAccountId(nil).Value()}

	if err := context.ShouldBindJSON(&json); err != nil {
		return NewResponseError(err)
	}

	handler.command = *create.NewCreateAccountCommand(
		json.Id,
		json.Name,
		json.Username,
		json.Password,
	)

	return nil
}

func (handler CreateAccountPostHandler) Handle(context *gin.Context) (*Response, *ResponseError) {
	if err := handler.validate(context); err != nil {
		return nil, err
	}

	if err := handler.commandBus.Dispatch(handler.command); err != nil {
		return nil, NewResponseError(err)
	}

	return NewResponse(
		http.StatusCreated,
		map[string]interface{}{
			"data": map[string]interface{}{
				"type": "accounts",
				"id":   handler.command.Data().(create.CreateAccountCommandData).Id,
			},
		},
	), nil
}
