package server

import (
	"credens/apps/http/handler"
	"credens/apps/http/handler/middleware"
	"credens/libs/shared/domain/bus"
	"github.com/defval/inject"
	"github.com/gin-gonic/gin"
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
