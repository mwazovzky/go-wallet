package database

import (
	"log"

	"database/sql"

	_ "modernc.org/sqlite"
)

func InitConnection(databasePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", databasePath)
	if err != nil {
		log.Println("failed to connect to database", err)
	}

	return db, err
}

func SetupDatabase(db *sql.DB) error {
	st := `CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		address TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := db.Exec(st)

	return err
}
