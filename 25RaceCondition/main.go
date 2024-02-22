package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("RACE CONDITION")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	Rmut := &sync.RWMutex{}

	// Read Mutex read lock only
	Rmut.RLock()
	var score = []int{0}
	Rmut.RUnlock()
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("One R")
		// mut here help us to prevent the race condition that any process coming to critical section any time , this can be like concept of OS
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)

	wg.Wait()
	Rmut.RLock()
	fmt.Println(score)
	Rmut.RUnlock()
}
