package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study")
	// presentTime := time.Now()
	// fmt.Println(presentTime.Local().Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.October, 5, 0, 0, 0, 0, time.Now().Local().Location())
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
