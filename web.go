package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// main starts an HTTP server listening on $PORT which dispatches to request handlers.
func main() {
	// wrap the hostedByHandler with logging middleware
	http.Handle("/", logRequestMiddleware(http.HandlerFunc(hostedByHandler)))
	port := os.Getenv("PORT")
	log.Printf("listening on %v...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// hostedByHandler writes "Hosted by $HOSTED_BY"
func hostedByHandler(w http.ResponseWriter, r *http.Request) {
	// print the string to the ResponseWriter
	hostname, _ := os.Hostname()
	log.Printf("METHOD: %s - URL: %s", r.Method, r.URL)
	fmt.Fprintf(w, "Hosted by %v\n", hostname)
}

// logRequestMiddleware writes out HTTP request information before passing to the next handler.
func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remote := r.RemoteAddr
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			remote = forwardedFor
		}
		log.Printf("%s %s %s", remote, r.Method, r.URL)
		// pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
