package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var JSON = `{
  "name": "Bob",
  "title": "Head chef",
  "contact": {
    "home": "123",
    "cell": "321"
  }
}`

func main() {
	var c map[string]any
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println("Name: ", c["name"])
	fmt.Println("home: ", c["contact"].(map[string]any)["home"])
	fmt.Println("cell: ", c["contact"].(map[string]any)["cell"])
}