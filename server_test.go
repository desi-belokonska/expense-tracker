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

		got := User{}

		err := json.NewDecoder(response.Body).Decode(&got)
		want := User{UserID: 1, FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"}

		if err != nil {
			t.Error("couldn't decode JSON into User")
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}

		assertContentType(t, response, contentTypeJSON)

	})
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	got := response.Result().Header.Get("content-type")

	if got != want {
		t.Errorf("response did not have correct content-type: got %q, want %q", got, want)
	}
}
