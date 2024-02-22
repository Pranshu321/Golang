package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func Proadder(values ...int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(add(2, 4))
	Proadder(2, 3, 4, 6, 4, 2)
}
