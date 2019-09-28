package server

import (
	"credens/apps/http/config"
	"credens/libs/shared/infrastructure/di"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(env config.Environment, container *di.Container) *Server {
	var mode = gin.ReleaseMode
	if env.Debug {
		switch env.Env {
		case "test", "testing":
			mode = gin.TestMode
		default:
			mode = gin.DebugMode
		}
	}
	gin.SetMode(mode)

	engine := gin.Default()

	addRoutes(engine, container)

	return &Server{engine}
}

func (server *Server) Run(port int) error {
	return server.engine.Run(fmt.Sprintf(":%d", port))
}
