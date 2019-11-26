package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	store, err := NewExpenseStoreSQL(prodDB)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server := NewExpenseServer(store)
	loggedServer := Logger(server)
	fmt.Printf("Listening on port: %v\t(http://localhost:%v)\n", port, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), loggedServer); err != nil {
		log.Fatalf("could not listen on port %v %v", port, err)
	}
}
