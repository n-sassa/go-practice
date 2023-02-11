package postgresql

import (
	"database/sql"
	"fmt"
)

const driverName = "postgres"

type PostgreSQLConnector struct {
	Conn *sql.DB
}

func NewPostgreSQLConnector() *PostgreSQLConnector {
	dsn := postgresqlConnInfo()
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return &PostgreSQLConnector{
		Conn: conn,
	}
}

func postgresqlConnInfo() string {
	dataSourceName := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		"postgres",
		"develop",
		"app_user",
		"password")

	return dataSourceName
}
