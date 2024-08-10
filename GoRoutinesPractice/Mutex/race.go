package race

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var Wg sync.WaitGroup
var mut sync.Mutex // pointer as , it take care that only one time one go routine can access the resource

func GetStatusCode(endpoint string) {
	defer Wg.Done() // Responsible for decrementing the wg
	data, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops there was an error", err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", data.StatusCode, endpoint)
	}
}

func main() {

	s := []string{
		"https://pranshuportfolio.netlify.app",
		"https://go.dev",
		"https://google.com",
	}

	for _, web := range s {
		go GetStatusCode(web)
		Wg.Add(1) // Only one go routine that is firing up
	}
	Wg.Wait()
}
