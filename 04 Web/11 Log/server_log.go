package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index Page!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Gopher!")
}

func startPanic(w http.ResponseWriter, r *http.Request) {
	panic("Starting a panic...")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/hello", hello).Methods("GET")
	r.HandleFunc("/panic", startPanic).Methods("GET")

	// Logging middleware
	// Gorilla handler logs out a standard apache format log
	logger := handlers.LoggingHandler(os.Stdout, r)

	log.Println("Starting server at http://localhost:8080/")
	log.Fatalln(http.ListenAndServe(":8080", middleware.PanicRecovery(logger)))
}
