package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Conversion
	// 1. Implicit Conversion
	// 2. Explicit Conversion

	println("Welcome to papa jones")
	in := bufio.NewReader(os.Stdin)
	print("Enter your name: ")
	name, _ := in.ReadString('\n')
	print("Hello " + name)
	println("Please give rating between 1 to 5")
	rating, _ := in.ReadString('\n')
	numRating, err := strconv.ParseInt(strings.TrimSpace(rating), 0, 64)

	if err != nil {
		fmt.Println("Error occured", err)
	} else {
		// numRating++
		fmt.Println("Thanks for rating ", numRating)
	}
}
