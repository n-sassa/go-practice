package clients

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "app_user"
	password = "password"
	dbname   = "demo"
)

type ClientBuilder struct{}

func NewClientBuilder() ClientBuilder {
	return ClientBuilder{}
}

func (b ClientBuilder) BuildSqlClient() *gorm.DB {
	dsn := "host=postgres port=5432 user=app_user password=password dbname=app_db sslmode=disable"
	sqlDB, err := sql.Open("pgx", dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}

	return db
}
