package main

import "fmt"

func main() {
	fmt.Println("Welcome World of array")

	// Array Declaration
	var fruitlist [4]string
	var vege = [3]string{"Potato", "Tomato", "Brinjal"}
	fmt.Println(vege)

	fruitlist[0] = "Apple"
	fruitlist[1] = "Orange"
	// fruitlist[0]= "Kiwi"
	fruitlist[3] = "Chiku"

	fmt.Println(fruitlist[0:2])
	fmt.Println(len(fruitlist))
}
