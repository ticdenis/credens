package sql_migration

import (
	"github.com/lopezator/migrator"
)

func New1569878041CreateAccountsTableMigration() *migrator.Migration {
	return createSimpleMigration(
		"Create accounts table",
		`
			CREATE TABLE accounts (
				id VARCHAR(36) NOT NULL,
				name VARCHAR(200) NOT NULL,
				username VARCHAR(200) NOT NULL,
				password VARCHAR(200) NOT NULL,
				PRIMARY KEY (id)
			) ENGINE=InnoDB;
		`,
	)
}
