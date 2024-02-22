package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	IsFree   bool     `json:"free"`
	Password string   `json:"-"`
	Rating   float32  `json:"rating"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Encoded of the JSON")
	// EncodeJSON()
	DecodeJson()
}

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJS", 99, true, "abc123", 2.35, []string{"web-dev", "js"}},
		{"MERN", 199, true, "abc123", 7.35, []string{"web-dev", "js", "mongoDB"}},
		{"MEAN", 499, true, "abc123", 9.35, []string{"web-dev", "angular"}},
		{"Astro", 499, true, "abc123", 9.35, nil},
	}

	// package this as the json
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // to give the indent
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	  {
		"coursename": "MERN",
        "price": 199,
        "free": true,
        "rating": 7.35,
		"tags": ["web-dev","js","mongoDB"]
	  }
	`)

	var lcoCourse course
	CheckValid := json.Valid(jsonDataFromWeb)
	if CheckValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Println(myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T \n", k, v, v)
	}
}
