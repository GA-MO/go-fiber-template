package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitSQLite() (*sql.DB, error) {

	var err error

	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Successfully connected to database")
	return DB, nil
}

func GetSQLite() (*sql.DB, error) {

	if DB == nil {
		db, err := InitSQLite()
		if err != nil {
			log.Fatal(err)
		}

		return db, nil
	}
	return DB, nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
