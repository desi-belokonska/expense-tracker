package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dir = "./db.json"

// ExpenseStoreSQL is an implementaton of ExpenseStore in sqlite3
type ExpenseStoreSQL struct {
	*sql.DB
}

// GetUser queries the database for a user and returns it if it's found, nil otherwise
func (es *ExpenseStoreSQL) GetUser(id int) *User {
	var user *User

	rows, err := es.Query("SELECT * FROM users WHERE user_id = ?", id)

	if err != nil {
		log.Fatalln(err)
	}

	var (
		userID    int64
		firstName string
		lastName  string
		email     string
	)

	for rows.Next() {
		if err = rows.Scan(&userID, &firstName, &lastName, &email); err != nil {
			log.Fatalln(err)
		}
		user = &User{userID, firstName, lastName, email}
	}

	return user
}

// NewExpenseStoreSQL returns a pointer to an initialized ExpenseStoreSQL
func NewExpenseStoreSQL() (*ExpenseStoreSQL, error) {
	e := ExpenseStoreSQL{}

	db, err := sql.Open("sqlite3", "./db/expense-tracker.db")
	if err != nil {
		return nil, err
	}

	e.DB = db

	return &e, nil
}
