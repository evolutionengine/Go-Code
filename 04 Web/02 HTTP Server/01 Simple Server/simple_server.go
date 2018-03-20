package main

import (
	"fmt"
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello, Welcome To Home Page!")
}

func main() {
	http.HandleFunc("/", home)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}