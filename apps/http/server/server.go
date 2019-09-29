package server

import (
	"credens/apps/http/config"
	"fmt"
	"github.com/defval/inject"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(env config.Environment, container *inject.Container) (*Server, error) {
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

	err := addRoutes(engine, container)
	if err != nil {
		return nil, err
	}

	return &Server{engine}, nil
}

func (server *Server) Run(port int) error {
	return server.engine.Run(fmt.Sprintf(":%d", port))
}
