package main

import (
	"encoding/json"
	"net/http"
)

// User represents a user
type User struct {
	UserID    int64
	FirstName string
	LastName  string
	Email     string
}

// ExpenseServer is an HTTP interface for Expense Tracking
func ExpenseServer(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(User{1, "Jane", "Doe", "jane.doe@example.com"})
}
