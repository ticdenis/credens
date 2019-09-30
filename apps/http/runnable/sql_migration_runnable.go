package runnable

import (
	"credens/apps/http/config"
	"credens/apps/http/migration/sql_migration"
	"github.com/defval/inject"
)

type SQLMigrationRunnable struct {
}

func NewSQLMigrationRunnable() *SQLMigrationRunnable {
	return &SQLMigrationRunnable{}
}

func (_ SQLMigrationRunnable) Run(container *inject.Container, env config.Environment) error {
	if env.Sql.Migrate {
		var sqlMigrator sql_migration.SQLMigrator
		if err := container.Extract(&sqlMigrator); err != nil {
			return err
		}
		return sqlMigrator.Run()
	}
	return nil
}
