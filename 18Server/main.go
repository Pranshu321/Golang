package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the GO Server")
	// PerformGetRequest()
	// PerformPostRequest()
	PerfromPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"
	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Content Length: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:3000/post"
	requestBody := strings.NewReader(`
		{"courseName": "Let's go with golang","price": 0 , "platform": "LCO"}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerfromPostFormRequest() {
	const myurl = "http://localhost:3000/postform"
	data := url.Values{}
	data.Add("firstname", "shubham")
	data.Add("lastname", "kumar")
	data.Add("email", "j7Hr5@example.com")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
