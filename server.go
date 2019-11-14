package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	GetUser(id int) User
}

// ExpenseServer is an HTTP interface for Expense Tracking
type ExpenseServer struct {
	store ExpenseStore
}

func (es *ExpenseServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userIDString := strings.TrimPrefix(r.URL.Path, "/users/")

	w.Header().Set("content-type", contentTypeJSON)
	enc := json.NewEncoder(w)

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		log.Printf("couldn't get UserID from URL path: '%v'", err)
		return
	}

	user := es.store.GetUser(userID)

	enc.Encode(user)

}
