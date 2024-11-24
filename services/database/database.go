package database

import (
	"log"

	"database/sql"

	_ "modernc.org/sqlite"
)

const databasePath = "./database/accounts.db"

func InitConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", databasePath)
	if err != nil {
		log.Println("failed to connect to database", err)
	}

	return db, err
}

func SetupDatabase(db *sql.DB) error {
	st := `CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		chain TEXT NOT NULL,
		address TEXT NOT NULL,
		password TEXT NOT NULL,
		UNIQUE(chain,address)
	);`

	_, err := db.Exec(st)

	return err
}
