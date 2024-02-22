package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")
	// If someone is listening to the channel , then only channel allows us to pass the value.
	myCh := make(chan int)
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		val, isChannelOpen := <-ch
		// This is designed because ther eis a problem in which when channel is closed and we printing value it will give zero and we passing zero to hte channel then also giving zero , so to tackle this we have determine the isChannelopen bool varaible.
		if isChannelOpen {
			fmt.Println(val)
			fmt.Println(isChannelOpen)
			fmt.Println(<-ch)
		}
		// fmt.Println(<-ch)
	}(myCh, wg, mut)

	go func(ch chan int, wg *sync.WaitGroup, m *sync.RWMutex) {
		myCh <- 5
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg, mut)

	wg.Wait()

	// fmt.Println(<-myCh)
	// myCh <- 5
}
