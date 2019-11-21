package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubExpenseStore struct {
	users  map[int]*User
	nextID int
}

func (s *StubExpenseStore) GetUser(id int) *User {
	user := s.users[id]
	return user
}

func TestGETUser(t *testing.T) {

	store := StubExpenseStore{map[int]*User{
		1: {1, "Jane", "Doe", "jane.doe@example.com"},
		2: {2, "Spencer", "White", "spencer.white@example.com"},
	}, 3}

	server := ExpenseServer{&store}

	t.Run("returns Jane's information as JSON", func(t *testing.T) {
		request := newGetUserRequest(1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getUserFromResponse(t, response)
		want := User{UserID: 1, FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"}

		assertUser(t, got, want)
		assertContentType(t, response, contentTypeJSON)
		assertStatusCode(t, response.Code, http.StatusOK)

	})

	t.Run("returns Spencer's information as JSON", func(t *testing.T) {
		request := newGetUserRequest(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getUserFromResponse(t, response)
		want := User{UserID: 2, FirstName: "Spencer", LastName: "White", Email: "spencer.white@example.com"}

		assertUser(t, got, want)
		assertContentType(t, response, contentTypeJSON)
		assertStatusCode(t, response.Code, http.StatusOK)

	})

	t.Run("returns 404 on missing user", func(t *testing.T) {
		request := newGetUserRequest(10)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getErrorFromResponse(t, response)
		want := ErrorJSON{NotFound, "Requested user 10 not found.", nil}

		assertError(t, got, want)
		assertContentType(t, response, contentTypeJSON)
		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func newGetUserRequest(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", id), nil)
	return req
}

func getUserFromResponse(t *testing.T, response *httptest.ResponseRecorder) User {
	t.Helper()

	user := User{}
	err := json.NewDecoder(response.Body).Decode(&user)

	if err != nil {
		t.Errorf("couldn't decode JSON into User: response body - %q, '%v'", response.Body, err)
	}

	return user
}

func getErrorFromResponse(t *testing.T, response *httptest.ResponseRecorder) ErrorJSON {
	t.Helper()

	errorJSON := ErrorJSON{}
	err := json.NewDecoder(response.Body).Decode(&errorJSON)

	if err != nil {
		t.Errorf("couldn't decode JSON into ErrorJSON: response body - %q, '%v'", response.Body, err)
	}

	return errorJSON
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	got := response.Result().Header.Get("content-type")

	if got != want {
		t.Errorf("response did not have correct content-type: got %q, want %q", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("response did not have correct status code: got %d, want %d", got, want)
	}
}

func assertUser(t *testing.T, got, want User) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrong User: got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, got, want ErrorJSON) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrong Error: got %q, want %q", got, want)
	}
}
