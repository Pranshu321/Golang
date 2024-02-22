package main

import (
	"fmt"
	"io"
	"net/http"
)

// const url = "https://lco.dev"
const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("WEB REQUESTS IN GOLANG")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of res is %T\n", res)
	fmt.Println("Content type: ", res.Header.Get("Content-Type"))

	defer res.Body.Close() // This is our response body, we have to close it
	data, err := io.ReadAll(res.Body)
	if err == nil {
		fmt.Println("Data is : \n", string(data))
	}

}
