package persistence

import (
	"credens/libs/accounts/domain"
	"database/sql"
	"errors"
)

type MysqlAccountRepository struct {
	db *sql.DB
}

func NewMysqlAccountRepository(db *sql.DB) *MysqlAccountRepository {
	return &MysqlAccountRepository{db: db}
}

func (repo MysqlAccountRepository) Add(account *domain.Account) (err error) {
	_, err = repo.createAccountsTableIfNotExists()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(`
		INSERT INTO accounts (id, name, username, password) VALUES (?, ?, ?, ?);
	`, account.Id().Value(), account.Name().Value(), account.Username().Value(), account.Password().Value())
	if err != nil {
		return err
	}

	return nil
}

func (repo MysqlAccountRepository) createAccountsTableIfNotExists() (sql.Result, error) {
	return repo.db.Exec(`
			CREATE TABLE IF NOT EXISTS accounts (
				id VARCHAR(36) NOT NULL,
				name VARCHAR(200) NOT NULL,
				username VARCHAR(200) NOT NULL,
				password VARCHAR(200) NOT NULL,
				PRIMARY KEY (id)
			) ENGINE=InnoDB;
		`)
}

func (repo MysqlAccountRepository) Search(id domain.AccountId) (*domain.Account, error) {
	_, err := repo.createAccountsTableIfNotExists()
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.Query("SELECT * FROM accounts WHERE id = ? LIMIT 1", id.Value())
	if err != nil {
		return nil, err
	}

	var data struct {
		id       string
		name     string
		username string
		password string
	}

	for rows.Next() {
		if err = rows.Scan(&data.id, &data.name, &data.username, &data.password); err != nil {
			return nil, err
		}

		return domain.NewAccount(
			domain.NewAccountId(data.id),
			domain.NewAccountName(data.name),
			domain.NewAccountUsername(data.username),
			domain.NewAccountPassword(data.password),
		), nil // TODO: Use a builder to avoid side effects like as record DomainEvents!
	}

	return nil, errors.New("Account not found!")
}
