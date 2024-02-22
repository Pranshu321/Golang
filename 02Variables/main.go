package main

import "fmt"

// := (Walarus operator) it is allowed to user this function or method but not in global scope.
// var jwtToken = 30000

// Constants
const Pi float32 = 3.14

func main() {
	// fmt.Println("hello")
	// var isString string = "dfdf"
	// var islogged bool = true
	// var username int = 3

	var smallVal uint8 = 127
	fmt.Printf("What is type of variable username: %T \n  ", smallVal)

	// var smallFloat float32 = 255.454353534434342
	var smallFloat float64 = 255.454353534434342
	fmt.Println(smallFloat)
	fmt.Printf("What is type of variable username: %T \n  ", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("What is type of variable username: %T \n  ", anotherVariable)

	// implicit type or way for declaring variables
	var web = "github.com"
	fmt.Println(web)
	// web = 4 // This is will give error as lexer already treat it as string so we can't assign int.

	// No var style
	numberOfUser := 30000
	print(numberOfUser)

	// Print Constant
	fmt.Println(Pi)
	fmt.Printf("What is type of variable username: %T \n  ", Pi)

}
