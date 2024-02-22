package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	// Pointers in GoLang

	// * sign tells that it is pointer
	// var one *int
	// When we dont put anyting in pointer it will be nil

	// & for address
	// add := &one

	num := 10
	var ptr *int = &num
	var doubleptr **int = &ptr

	// Operation
	*ptr = *ptr + 2
	fmt.Println(num)

	// fmt.Println(*ptr)
	fmt.Println(**doubleptr)
}
