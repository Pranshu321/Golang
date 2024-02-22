package main

import (
	"fmt"
)

// When we use the defer the function will be executed at the end of the program (means just before the end } braces)

// The defer execute will be in reverse order means LIFO

func myDefer() {
	fmt.Println("This is my defer")
	for i := 6; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("1")
	myDefer()
}

// 1
// 9 ,8 , 7, 6
// 4 , 3 , 2
