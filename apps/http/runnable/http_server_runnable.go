package runnable

import (
	"credens/apps/http/config"
	"credens/apps/http/server"
	"github.com/defval/inject"
)

type HttpServerRunnable struct {
}

func NewHttpServerRunnable() *HttpServerRunnable {
	return &HttpServerRunnable{}
}

func (_ HttpServerRunnable) Run(container *inject.Container, env config.Environment) error {
	svc, err := server.NewServer(env, container)
	if err != nil {
		return err
	}
	return svc.Run(env.Port)
}
