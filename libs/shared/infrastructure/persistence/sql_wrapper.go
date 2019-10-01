package db

import (
	"database/sql"
)

type (
	SQLDb interface {
		Run() error
		DB() *sql.DB
		open() error
		close() error
	}

	SQLDBWrapper struct {
		driverName string
		url        string
		db         *sql.DB
	}
)

func (wrapper *SQLDBWrapper) Run() error {
	if err := wrapper.open(); err != nil {
		return err
	}
	// wrapper.close()
	return nil
}

func (wrapper *SQLDBWrapper) open() error {
	if wrapper.db != nil {
		return nil
	}

	db, err := sql.Open(wrapper.driverName, wrapper.url)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	wrapper.db = db

	return nil
}

func (wrapper *SQLDBWrapper) close() error {
	if wrapper.db != nil {
		return wrapper.db.Close()
	}
	return nil
}

func (wrapper *SQLDBWrapper) DB() *sql.DB {
	return wrapper.db
}

func NewSQLWrapper(driverName, url string) *SQLDBWrapper {
	return &SQLDBWrapper{driverName: driverName, url: url}
}
