package persistence

import (
	"credens/libs/accounts/domain"
	db "credens/libs/shared/infrastructure/persistence"
	"database/sql"
	"errors"
)

type MysqlAccountRepository struct {
	db *sql.DB
}

func NewMysqlAccountRepository(sql db.SQLDb) *MysqlAccountRepository {
	return &MysqlAccountRepository{db: sql.DB()}
}

func (repo MysqlAccountRepository) Add(account *domain.Account) error {
	_, err := repo.db.Exec(`
		INSERT INTO accounts (id, name, username, password) VALUES (?, ?, ?, ?);
	`, account.Id().Value(), account.Name().Value(), account.Username().Value(), account.Password().Value())
	if err != nil {
		return err
	}

	return nil
}

func (repo MysqlAccountRepository) Search(id domain.AccountId) (*domain.Account, error) {
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
