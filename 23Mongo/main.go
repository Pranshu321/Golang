package main

import (
	"fmt"
	"log"
	router "mongo/routers"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	r := router.Router()
	fmt.Println("Server is Starting ....")
	// http.ListenAndServe(":3000", r);
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Server is Started 😔 ....")
}
