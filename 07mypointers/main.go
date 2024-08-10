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
	ptr := &num
	dptr := &ptr
	fmt.Println("Value of num is ", num)
	*ptr = *ptr + 2
	fmt.Println("Value of num is ", num)

	fmt.Println("Value of [[ptr]] is ", **dptr)
}
