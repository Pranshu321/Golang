package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is map")
	// Define mao
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["TS"] = "Typescript"

	// delete
	delete(languages, "RB")

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	fmt.Println(languages)

}
