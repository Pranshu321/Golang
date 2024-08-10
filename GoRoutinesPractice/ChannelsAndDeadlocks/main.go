package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and Deadlocks")

	// Creating a channel
	ch := make(chan int, 2)
	// ch2 := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// without the listener, the program will deadlock, will not run

	// fmt.Println(<-ch)
	// ch <- 5

	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, isChannelOpen := <-ch // if the channel is closed, the value will be 0 and the channel will be closed. This is imp because we are not able to distinguish between the two cases of channel being closed and the value being 0 as in both cases the value will be 0

		fmt.Println(val, isChannelOpen)
		// fmt.Println(<-ch)
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		// ch <- 0
		close(ch)
		// ch <- 6
		defer wg.Done()
	}(ch, wg)

	wg.Wait()

}
