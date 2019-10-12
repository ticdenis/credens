package runnable

import (
	"credens/apps/amqp-worker/config"
	"github.com/defval/inject"
)

type Runnable interface {
	Run(container *inject.Container, env config.Environment) error
}
