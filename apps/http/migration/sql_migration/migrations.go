package sql_migration

import (
	db "credens/libs/shared/infrastructure/persistence"
	"database/sql"
	"github.com/lopezator/migrator"
)

func createSimpleMigration(name, query string, args ...interface{}) *migrator.Migration {
	return &migrator.Migration{
		Name: name,
		Func: func(tx *sql.Tx) error {
			if _, err := tx.Exec(query, args...); err != nil {
				return err
			}
			return nil
		},
	}
}

type (
	SQLMigrator interface {
		Run() error
	}

	SQLMigratorWrapper struct {
		sql db.SQLDb
	}
)

func NewSQLMigratorWrapper(sql db.SQLDb) *SQLMigratorWrapper {
	return &SQLMigratorWrapper{sql: sql}
}

func (m *SQLMigratorWrapper) Run() error {
	sqlMigrator, err := migrator.New(m.getMigrations())
	if err != nil {
		return err
	}
	return sqlMigrator.Migrate(m.sql.DB())
}

func (m *SQLMigratorWrapper) getMigrations() migrator.Option {
	return migrator.Migrations(
		New1569878041CreateAccountsTableMigration(),
	)
}
