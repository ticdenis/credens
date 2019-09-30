package runnable

import (
	"credens/apps/http/config"
	db "credens/libs/shared/infrastructure/persistence"
	"github.com/defval/inject"
)

type SQLDatabaseRunnable struct {
}

func NewSQLDatabaseRunnable() *SQLDatabaseRunnable {
	return &SQLDatabaseRunnable{}
}

func (_ SQLDatabaseRunnable) Run(container *inject.Container, env config.Environment) error {
	var sqlDB db.SQLDb
	if err := container.Extract(&sqlDB); err != nil {
		return err
	}
	return sqlDB.Run();
}
