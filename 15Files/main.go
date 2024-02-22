package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World! For File")
	in := bufio.NewReader(os.Stdin)
	content, _ := in.ReadString('\n')
	content = strings.TrimSpace(content)
	ReadFile("test.txt", "./15Files/test.txt")

	// fmt.Println("Content is ", content)

	file, err := os.Create("./15Files/test.txt")
	if err != nil {
		fmt.Println("File Not Created !!")
		panic(err)
	} else {
		length, err := io.WriteString(file, content)
		if err != nil {
			fmt.Println("Content Not Written !!")
			panic(err)
		}
		fmt.Println(length)
	}
	defer file.Close()
}

func ReadFile(fileName string, path string) {
	databyte, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("File Not Present")
		panic(err)
	}
	fmt.Println("This is File Content : ", string(databyte))
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
