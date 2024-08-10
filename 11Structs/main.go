package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type person struct {
	Name   string
	Age    int
	Email  string
	status bool
}

// Methofs

// Getter Functions
func (p person) GetStatus() {
	fmt.Println("is user active: ", p.status)
}

func (p *person) SetAge(newAge int) {
	p.Age = newAge
}

// Setter Functions

func main() {
	pranshu := person{
		Name:   "Pranshu",
		Age:    20,
		Email:  "p@p.com",
		status: false,
	}
	// arr := []person{}
	// arr[0] = person{"John", 30, "j@j.com"}
	// arr[1] = person{"JD", 30, "j@j.com"}
	// arr[2] = person{"df", 30, "j@j.com"}
	// arr[3] = person{"dfdf", 30, "j@j.com"}

	// fmt.Println(arr[0])

	fmt.Println(pranshu)
	pranshu.GetStatus()
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Enter new age: ")
	val, _ := in.ReadString('\n')
	intval, _ := strconv.Atoi(strings.TrimSpace(val))
	pranshu.SetAge(intval)
	fmt.Println(pranshu)

	// fmt.Printf("pranshu details: %+v\n", pranshu)
	// fmt.Println(pranshu.Name)
	// fmt.Println(pranshu.Age)
	// fmt.Println(pranshu.Email)

}
