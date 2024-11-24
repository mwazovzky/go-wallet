package repository

import (
	"database/sql"
	"log"
)

type Account struct {
	ID       int64
	Chain    string
	Address  string
	Password string
}

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) Fetch() ([]Account, error) {
	query := `SELECT id, address, password FROM accounts;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []Account{}

	for rows.Next() {
		var account Account
		err := rows.Scan(&account.ID, &account.Address, &account.Password)
		if err != nil {
			log.Fatal("failed to scan query data", err)
		}
		data = append(data, account)
	}

	return data, nil
}

func (r *AccountRepository) Find(address string) (*Account, error) {
	query := `SELECT id, address, password FROM accounts WHERE address=? LIMIT 1;`

	rows, err := r.db.Query(query, address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var account Account

	rows.Next()
	err = rows.Scan(&account.ID, &account.Address, &account.Password)
	if err != nil {
		log.Fatal("failed to scan query data", err)
	}

	return &account, nil
}

func (r *AccountRepository) Create(account Account) error {
	statement := `INSERT INTO accounts (address, password) VALUES (?,?)`

	_, err := r.db.Exec(statement, account.Address, account.Password)

	return err
}
