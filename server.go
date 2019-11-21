package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const contentTypeJSON = "application/json"

// User represents a user
type User struct {
	UserID    int64  `json:"userID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// ExpenseStore stores all expense related information about the users
type ExpenseStore interface {
	GetUser(id int) *User
}

// ExpenseServer is an HTTP interface for Expense Tracking
type ExpenseServer struct {
	store ExpenseStore
	http.Handler
}

// NewExpenseServer returs a pointer to a new ExpenseServer with the given store
func NewExpenseServer(store ExpenseStore) *ExpenseServer {
	e := new(ExpenseServer)

	e.store = store

	router := mux.NewRouter()
	router.HandleFunc("/users/{userID}", http.HandlerFunc(e.usersHandler))

	e.Handler = router

	return e
}

func (es *ExpenseServer) usersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars["userID"]
	userID, err := strconv.Atoi(userIDString)

	w.Header().Set("content-type", contentTypeJSON)
	enc := json.NewEncoder(w)

	if err != nil {
		log.Printf("couldn't get UserID from URL path: '%v'", err)
		w.WriteHeader(http.StatusNotFound)
		errorJSON := CreateErrorNotFound(fmt.Sprintf("Couldn't get UserID from URL path: %v", userIDString))
		enc.Encode(errorJSON)
		return
	}

	user := es.store.GetUser(userID)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		errorJSON := CreateErrorNotFound(fmt.Sprintf("Requested user %v not found.", userID))
		enc.Encode(errorJSON)
		return
	}

	w.WriteHeader(http.StatusOK)
	enc.Encode(user)

}
