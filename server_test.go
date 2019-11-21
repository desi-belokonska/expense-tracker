package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// ==== Mocks/Stubs/Spies

type StubExpenseStore struct {
	users  map[int]*User
	nextID int
}

func (s *StubExpenseStore) GetUser(id int) *User {
	user := s.users[id]
	return user
}

// ==== Tests

func TestGETUserSuccess(t *testing.T) {
	testCases := []struct {
		desc       string
		userID     int
		wantedUser User
	}{
		{
			desc:       "returns Jane's information as JSON",
			userID:     1,
			wantedUser: User{UserID: 1, FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"},
		},
		{
			desc:       "returns Spencer's information as JSON",
			userID:     2,
			wantedUser: User{UserID: 2, FirstName: "Spencer", LastName: "White", Email: "spencer.white@example.com"},
		},
	}
	store := StubExpenseStore{map[int]*User{
		1: {1, "Jane", "Doe", "jane.doe@example.com"},
		2: {2, "Spencer", "White", "spencer.white@example.com"},
	}, 3}

	server := NewExpenseServer(&store)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			request := newGetUserRequest(tC.userID)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			got := getUserFromResponse(t, response)

			assertUser(t, got, tC.wantedUser)
			assertContentType(t, response, contentTypeJSON)
			assertStatusCode(t, response.Code, http.StatusOK)
		})
	}
}

func TestGETUserFailure(t *testing.T) {

	store := StubExpenseStore{map[int]*User{
		1: {1, "Jane", "Doe", "jane.doe@example.com"},
		2: {2, "Spencer", "White", "spencer.white@example.com"},
	}, 3}

	server := NewExpenseServer(&store)

	t.Run("returns 404 and an error on missing user", func(t *testing.T) {
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

// ==== heplers

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

// ==== assert helpers

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	got := response.Result().Header.Get("content-type")

	if got != want {
		t.Errorf("response did not have correct content-type: got %q, want %q", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()

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
