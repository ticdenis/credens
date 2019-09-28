package server

import (
	"credens/apps/http/config"
	"credens/apps/http/handler"
	"credens/apps/http/handler/middleware"
	"credens/libs/shared/domain/bus"
	"credens/libs/shared/infrastructure/di"
	"github.com/gin-gonic/gin"
)

func addRoutes(server *gin.Engine, container *di.Container) {
	handlerMiddleware := middleware.NewJSONHandlerMiddleware()

	commandBus := container.Get(config.CommandBusKey).(bus.CommandBus)
	queryBus := container.Get(config.QueryBusKey).(bus.QueryBus)

	server.GET("/healthz", handlerMiddleware.Handle(handler.NewHealthzGetHandler()))

	server.GET("/accounts/:id", handlerMiddleware.Handle(handler.NewReadAccountGetHandler(queryBus)))

	server.POST("/accounts", handlerMiddleware.Handle(handler.NewCreateAccountPostHandler(commandBus)))
}
