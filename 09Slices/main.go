package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	// slices

	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	// How to add data
	fruitlist = append(fruitlist, "Mango", "Banana")

	// Slicing
	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	// highScores
	highScores := make([]int, 4)
	highScores[0] = 98
	highScores[1] = 76
	highScores[2] = 56
	highScores[3] = 5

	// Append will re-allocate the memory when we use append
	highScores = append(highScores, 1, 2, 3, 8, 6, 7)

	sort.Slice(highScores, func(i, j int) bool { return highScores[i] < highScores[j] })

	if sort.IntsAreSorted(highScores) {
		print("True\n")
	}
	fmt.Println(highScores)

	// HOW TO REMOVE A VALUE FROM SLICE FROM INDEX
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	index := 2

	// what ... means in Golang
	// Suppose you have two slices of numbers, and you want to combine them. The ... notation can help you here too.
	// The slice2... notation tells Go to take each element of slice2 and append it to slice1.
	//  ... is for variadic arguments
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
