package main

import (
	"bufio"
	"fmt"
	"os"
)

func Loops() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	arr := []string{
		"Pranshu",
		"Welcome",
		"Riya",
		"Ishh",
		"Isha",
		"Meghwanshi"}

	for idx, value := range arr {
		if value == "Ishh" {
			goto lco
		}
		// if idx == 4 {
		// 	break
		// }
		fmt.Println(idx, value)
	}
lco:
	fmt.Println("Hello ishh")
}

func main() {
	welcome := "Welcome Here"
	fmt.Println(welcome)
	Loops()
	// panic("Break")

	fmt.Println("Enter the rating for our Pizza:")

	// comma ok syntax or Error ok syntax

	// When you don't care about the thing you can put underscore
	// input, _ := read.ReadString('\n')
	read := bufio.NewReader(os.Stdin)
	input, _ := read.ReadString('\n')
	// print(err.Error())
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of rating %T", input)
}
