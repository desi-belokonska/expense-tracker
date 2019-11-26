package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gchaincl/dotsql"
	_ "github.com/mattn/go-sqlite3"
)

const prodDB = "expense-tracker"
const testDB = "expense-tracker-test"

// ExpenseStoreSQL is an implementaton of ExpenseStore in sqlite3
type ExpenseStoreSQL struct {
	queries *dotsql.DotSql
	*sql.DB
}

// GetUser queries the database for a user and returns it if it's found, nil otherwise
func (es *ExpenseStoreSQL) GetUser(id int64) *User {
	var user *User

	rows, err := es.queries.Query(es, "get-user-by-id", id)

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

// GetUsers queries the database for all users and returns them if it's found, nil otherwise
func (es *ExpenseStoreSQL) GetUsers() *SliceResponse {
	ur := &SliceResponse{Users: []User{}}

	rows, err := es.queries.Query(es, "get-users")

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
		ur.Users = append(ur.Users, User{userID, firstName, lastName, email})
	}

	return ur
}

// NewUser adds a user to the DB and returns the new information as a pointer to a User
func (es *ExpenseStoreSQL) CreateUser(user User) *User {
	res, err := es.queries.Exec(es, "create-user", user.FirstName, user.LastName, user.Email)

	if err != nil {
		log.Fatalln(err)
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		log.Fatalln(err)
	}
	addedUser := es.GetUser(lastID)

	return addedUser

}

// ClearStore clears all info inside the store
func (es *ExpenseStoreSQL) ClearStore() {
	_, err := es.queries.Exec(es, "drop-users-table")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = es.queries.Exec(es, "create-users-table")
	if err != nil {
		log.Fatalln(err)
	}
}

// NewExpenseStoreSQL returns a pointer to an initialized ExpenseStoreSQL
func NewExpenseStoreSQL(name string) (*ExpenseStoreSQL, error) {
	e := ExpenseStoreSQL{}

	location := fmt.Sprintf("./db/%s.db", name)

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}

	dot, err := dotsql.LoadFromFile("db/scripts/queries.sql")

	if err != nil {
		return nil, err
	}

	e.DB = db
	e.queries = dot

	return &e, nil
}
