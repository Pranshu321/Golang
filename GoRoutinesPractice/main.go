package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func greeter(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Millisecond) // This basically adding waititng , so go routine can be executed
		fmt.Println(s)
	}
}

func GetStatusCode(endpoint string) {
	defer Wg.Done() // Responsible for decrementing the wg
	data, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops there was an error", err)
	}
	fmt.Printf("%d status code for %s\n", data.StatusCode, endpoint)
}

var Wg sync.WaitGroup // This is a global variable and pointer to the WaitGroup

func main() {

	/*
		go greeter("Hello") // never waits for the go routine to finish
		greeter("World")
	*/

	s := []string{
		"https://pranshuportfolio.netlify.app",
		"https://go.dev",
		"https://google.com",
		"https://prnsh.com",
	}

	for _, web := range s {
		go GetStatusCode(web)
		Wg.Add(1) // Only one go routine that is firing up
	}
	Wg.Wait() // Only responsible for waiting for all the go routines
}
