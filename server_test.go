package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETUser(t *testing.T) {
	t.Run("returns user as JSON", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		response := httptest.NewRecorder()

		ExpenseServer(response, request)

		got := User{}

		err := json.NewDecoder(response.Body).Decode(&got)
		want := User{UserID: 1, FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"}

		if err != nil {
			t.Errorf("couldn't decode into User")
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
