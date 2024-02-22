package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("World")
	s := []string{
		"https://pranshuportfolio.netlify.app",
		"https://go.dev",
		"https://google.com",
	}
	for _, web := range s {
		go getStatusCode(web)
		wg.Add(1) // Only one go routine that is firing up
	}
	wg.Wait() // Only responsible for waiting for all the go routines
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done() // Responsible for decrementing the wg
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d statuscode for the %s \n", res.StatusCode, endpoint)
}
