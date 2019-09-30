package runnable

import (
	"credens/apps/http/config"
	"github.com/defval/inject"
)

type Runnable interface {
	Run(container *inject.Container, env config.Environment) error
}
