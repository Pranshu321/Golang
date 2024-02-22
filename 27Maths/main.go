// package maths
package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
	// "math/rand"
)

func main() {
	fmt.Println("Maths in GOLANG")

	var num int = 5
	var myfloat float64 = 5.3

	fmt.Println(num + int(myfloat))

	// random
	// rand.Seed(time.Now().UnixNano())
	// ran := rand.Intn(5) + 1
	// fmt.Println(ran)

	// random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNumber)
}
