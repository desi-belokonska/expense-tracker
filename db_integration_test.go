package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddingUsersAndRetrievingThem(t *testing.T) {
	store, err := NewExpenseStoreSQL(testDB)

	if err != nil {
		t.Errorf("Problem initializing DB: %v", err)
	}

	store.ClearStore()
	server := NewExpenseServer(store)

	user1JSON := []byte(`{
		"firstName": "Ellen",
		"email": "ellie@example.com"
	}`)

	user2JSON := []byte(`{
		"firstName": "Thomas",
		"lastName": "Henry",
		"email": "thomas.henry@example.com"
	}`)

	server.ServeHTTP(httptest.NewRecorder(), newCreateUserRequest(user1JSON))
	server.ServeHTTP(httptest.NewRecorder(), newCreateUserRequest(user2JSON))

	t.Run("get 1st user", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetUserRequest(1))

		got := getUserFromResponse(t, response)
		want := User{1, "Ellen", "", "ellie@example.com"}

		assertStatusCode(t, response.Code, http.StatusOK)
		assertUser(t, got, want)
	})

	t.Run("get 2nd user", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetUserRequest(2))

		got := getUserFromResponse(t, response)
		want := User{2, "Thomas", "Henry", "thomas.henry@example.com"}

		assertStatusCode(t, response.Code, http.StatusOK)
		assertUser(t, got, want)
	})

	t.Run("get all users", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetUsersRequest())

		got := getUsersFromResponse(t, response)
		want := []User{
			{1, "Ellen", "", "ellie@example.com"},
			{2, "Thomas", "Henry", "thomas.henry@example.com"},
		}

		assertStatusCode(t, response.Code, http.StatusOK)
		assertUsers(t, got, want)
	})
}
