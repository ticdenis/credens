package sql_migration

import (
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

func getMigrations() migrator.Option {
	return migrator.Migrations(
		New1569878041CreateAccountsTableMigration(),
	)
}

func Migrate(db *sql.DB) error {
	sqlMigrator, err := migrator.New(getMigrations())
	if err != nil {
		return err
	}

	return sqlMigrator.Migrate(db)
}
