package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.w3schools.com/html/tryit.asp?filename=tryhtml_editor"

// const myurl = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Handling the URL")
	fmt.Println(myurl)
	// parsing
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Hostname(), result.Scheme, result.Path)
	// fmt.Println(result.RawQuery) // params
	// fmt.Println(result.Port())

	// queryParams
	qparams := result.Query() // This will give map
	fmt.Println(qparams.Get("filename"))

	// Partsofurl
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh"}

	anotherurl := partsOfUrl.String()
	fmt.Println(anotherurl)
}
