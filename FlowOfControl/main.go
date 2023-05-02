package main

import (
	"log"
	"net/http"
)

func middlewareA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//this is executed on the way down to the handler

		log.Println("Executing middleware A")

		next.ServeHTTP(w, r)
		log.Println("Executing middlewear a again")

	})
}

func middlewareB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//this is executed on the way down to the handler

		log.Println("Executing middleware B")

		if r.URL.Path == "/cherry" {
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middlewear b again")

	})
}

//create handler function

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from UB \n"))
	w.Write([]byte("Hello from UB \n"))
}

func ourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing he handler")
	w.Write([]byte("Carrots \n"))
	if r.URL.Path == "/cherry" {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", middlewareB(middlewareA(http.HandlerFunc(ourHandler))))

	mux.Handle("/home", http.HandlerFunc(home))

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
