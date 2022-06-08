package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var JSON = `{
  "name": "Bob",
  "title": "Head chef",
  "Contact": {
      "Home": "123",
      "Cell":  "321"
  }
}`

func main() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println(c)
}