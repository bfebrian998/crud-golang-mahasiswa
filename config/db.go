package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConfigDatabases() *sql.DB {
	conStr := "user=postgres password=mysecretpassword dbname=crud_db host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal("failed to open database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("could not connect to the database", err)
	}

	fmt.Println("Successfully connected to the database")
	return db
}

func InitDb() (*sql.DB, error) {
	db := ConfigDatabases()
	if db == nil {

		return nil, fmt.Errorf("failed to initialize database")
	}
	return db, nil

}
