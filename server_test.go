package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETUser(t *testing.T) {
	t.Run("returns Jane's information as JSON", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		response := httptest.NewRecorder()

		ExpenseServer(response, request)

		got := getUserFromResponse(t, response)
		want := User{UserID: 1, FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"}

		assertUser(t, got, want)
		assertContentType(t, response, contentTypeJSON)

	})

	t.Run("returns Spencer's information as JSON", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/2", nil)
		response := httptest.NewRecorder()

		ExpenseServer(response, request)

		got := getUserFromResponse(t, response)
		want := User{UserID: 2, FirstName: "Spencer", LastName: "White", Email: "spencer.white@example.com"}

		assertUser(t, got, want)
		assertContentType(t, response, contentTypeJSON)

	})
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

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	got := response.Result().Header.Get("content-type")

	if got != want {
		t.Errorf("response did not have correct content-type: got %q, want %q", got, want)
	}
}

func assertUser(t *testing.T, got, want User) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
