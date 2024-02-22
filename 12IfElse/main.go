package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetRandom() {
	fmt.Println("Switch Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		{
			fmt.Println("Value of dice is 1")
		}
	case 2:
		fmt.Println("Value of dice is 2")
		fallthrough
	case 3:
		fmt.Println("Value of dice is 3")
		fallthrough // This basically allow us to check futher cases or execute futher casess
	case 4:
		{
			fmt.Println("Value of dice is 4")
		}
	case 5:
		{
			fmt.Println("Value of dice is 5")
		}
	case 6:
		{
			fmt.Println("Value of dice is 6")
		}
	default:
		{
			fmt.Print("Out of range")
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	num, _ := in.ReadString('\n')
	numRating, _ := strconv.ParseInt(strings.TrimSpace(num), 0, 64)

	if numRating < 0 {
		fmt.Println("Negative")
	} else if numRating == 0 {
		panic("Zero is neither positive nor negative")
	} else {
		fmt.Println("Positive")
	}
	GetRandom()

}
