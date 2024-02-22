package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the modImportance")
	greet := greeter()
	fmt.Println(greet)
	e := mux.NewRouter()
	e.HandleFunc("/", serve)
	http.ListenAndServe(":3000", e)
}

func greeter() string {
	return "Hello there"
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my server</h1>"))
}
