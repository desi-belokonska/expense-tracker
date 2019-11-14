package main

import (
	"log"
	"net/http"
)

func main() {
	store, err := NewExpenseStoreSQL()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	server := ExpenseServer{store}
	if err := http.ListenAndServe(":5000", &server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
