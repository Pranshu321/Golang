package gorouties

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")
	var score = make([]int, 1)

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
