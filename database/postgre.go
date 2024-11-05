package database

import (
	"20241105/class/2/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func PgConnect() (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost", config.DbUser, config.DbPassword, config.DbName))
}

func DbOpen() *sql.DB {
	db, err := PgConnect()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
