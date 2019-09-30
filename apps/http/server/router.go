package server

import (
	"github.com/defval/inject"
	"github.com/gin-gonic/gin"

	"credens/libs/shared/domain/bus"

	"credens/apps/http/handler"
	"credens/apps/http/handler/middleware"
)

func addRoutes(server *gin.Engine, container *inject.Container) error {
	handlerMiddleware := middleware.NewJSONHandlerMiddleware()

	var commandBus bus.CommandBus
	if err := container.Extract(&commandBus); err != nil {
		return err
	}

	var queryBus bus.QueryBus
	if err := container.Extract(&queryBus); err != nil {
		return err
	}

	server.GET("/healthz", handlerMiddleware.Handle(handler.NewHealthzGetHandler()))

	server.GET("/accounts/:id", handlerMiddleware.Handle(handler.NewReadAccountGetHandler(queryBus)))

	server.POST("/accounts", handlerMiddleware.Handle(handler.NewCreateAccountPostHandler(commandBus)))

	return nil
}
