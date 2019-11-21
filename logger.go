package main

import (
	"log"
	"net/http"
	"time"
)

// Logger writes request information to standard output
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
