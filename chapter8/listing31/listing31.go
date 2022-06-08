package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]any)
	c["name"] = "gopher"
	c["title"] = "rapscallion"
	c["contact"] = map[string]any{
		"home": "123",
		"cell": "321",
	}

	rawJson, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println(rawJson)
}