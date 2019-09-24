package main

import (
	"credens/apps/http/handler"
	"credens/libs/shared/domain/bus"
	"credens/libs/shared/infrastructure/di"
	"github.com/gin-gonic/gin"
)

func AddRoutes(server *gin.Engine, container *di.Container) {
	server.GET("/healthz", handler.JSONHandler(handler.NewHealthzGetHandler()))

	server.GET("/accounts/:id", handler.JSONHandler(handler.NewReadAccountGetHandler(
		container.Get(QueryBusKey).(bus.QueryBus),
	)))

	server.POST("/accounts", handler.JSONHandler(handler.NewCreateAccountPostHandler(
		container.Get(CommandBusKey).(bus.CommandBus),
	)))
}
