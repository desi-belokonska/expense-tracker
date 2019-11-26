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

// SliceResponse contains a slice of User to return as JSON
// SliceResponse needs to be a struct because returning arrays in JSON can be problematic
// (see: https://haacked.com/archive/2009/06/25/json-hijacking.aspx/)
type SliceResponse struct {
	Users []User `json:"users"`
}

// ExpenseStore stores all expense related information about the users
type ExpenseStore interface {
	GetUser(id int64) *User
	GetUsers() *SliceResponse
	CreateUser(user User) *User
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

	router.HandleFunc("/users", e.usersHandler)
	router.HandleFunc("/users/{userID}", e.userHandler)

	e.Handler = router

	return e
}
func (es *ExpenseServer) usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		es.getUsers(w, r)
	case http.MethodPost:
		es.postUsers(w, r)
	}
}

func (es *ExpenseServer) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", contentTypeJSON)
	enc := json.NewEncoder(w)

	userResponse := es.store.GetUsers()

	w.WriteHeader(http.StatusOK)
	enc.Encode(userResponse)
}

func (es *ExpenseServer) postUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", contentTypeJSON)
	enc := json.NewEncoder(w)
	dec := json.NewDecoder(r.Body)

	var newUser User

	err := dec.Decode(&newUser)

	if err != nil {
		log.Printf("problem with POST /users: '%v'", err)
		w.WriteHeader(http.StatusNotFound)
		errorJSON := CreateErrorInvalidInput("Couldn't create a new user from input.", nil)
		enc.Encode(errorJSON)
		return
	}

	user := es.store.CreateUser(newUser)

	w.WriteHeader(http.StatusOK)
	enc.Encode(user)
}

func (es *ExpenseServer) userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", contentTypeJSON)
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	userIDString := vars["userID"]

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		log.Printf("couldn't get UserID from URL path: '%v'", err)
		w.WriteHeader(http.StatusNotFound)
		errorJSON := CreateErrorNotFound(fmt.Sprintf("Couldn't get UserID from URL path: %v", userIDString))
		enc.Encode(errorJSON)
		return
	}

	user := es.store.GetUser(int64(userID))

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		errorJSON := CreateErrorNotFound(fmt.Sprintf("Requested user %v not found.", userID))
		enc.Encode(errorJSON)
		return
	}

	w.WriteHeader(http.StatusOK)
	enc.Encode(user)

}
