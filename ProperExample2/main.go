package main

import (
	"fmt"
	"log"
	"net/http"
)

func logUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr

		forwardedFor := r.Header.Get("X-Forwarded-For")
		if forwardedFor != "" {
			ip = forwardedFor
		}

		fmt.Fprintf(w, "IP address: %s \n", ip)

		next.ServeHTTP(w, r)

	})
}

func logUrlPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/cancel" {
			next.ServeHTTP(w, r)
			return
		}

		fmt.Fprintf(w, "Requested path: %s \n", r.URL.Path)
		next.ServeHTTP(w, r)

	})
}

func ourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing he handler")
	w.Write([]byte("Information Logged \n"))

}

func main() {
	mux := http.NewServeMux()

	// mux.Handle("/", logUser(http.HandlerFunc(ourHandler)))
	mux.Handle("/", logUrlPath(logUser(http.HandlerFunc(ourHandler))))

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
