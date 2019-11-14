package main

import (
	"encoding/json"
	"net/http"
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

// ExpenseServer is an HTTP interface for Expense Tracking
func ExpenseServer(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/users/")

	w.Header().Set("content-type", contentTypeJSON)
	enc := json.NewEncoder(w)

	if userID == "1" {
		enc.Encode(User{1, "Jane", "Doe", "jane.doe@example.com"})
		return
	}

	if userID == "2" {
		enc.Encode(User{2, "Spencer", "White", "spencer.white@example.com"})
	}

}
