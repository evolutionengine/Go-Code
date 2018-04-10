package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Home Page!")
}

func pages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fmt.Fprintf(w, "Page No: %v\n", pageID)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/pages/{id:[0-9]+}/", pages)
	log.Fatalln(http.ListenAndServe(":8083", mux))
}



