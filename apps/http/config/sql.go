package config

import (
	db "credens/libs/shared/infrastructure/persistence"
	"fmt"
	"github.com/go-sql-driver/mysql" // It loads MySQL driver
)

func NewMySQLDBWrapper(env Environment) *db.SQLDBWrapper {
	cnf := mysql.Config{
		User:                 env.Sql.User,
		Passwd:               env.Sql.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", env.Sql.Host, env.Sql.Port),
		DBName:               env.Sql.Database,
		AllowNativePasswords: true,
	}
	return db.NewSQLWrapper(env.Sql.Driver, cnf.FormatDSN())
}
