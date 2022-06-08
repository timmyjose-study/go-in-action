package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	result struct {
		API         string `json:"API"`
		Description string `json:"Description"`
		Auth        string `json:"Auth"`
		Https       bool   `json:HTTPS`
		Cors        string `json:"Cors"`
		Link        string `json:"Link"`
		Category    string `json:"Category"`
	}

	response struct {
		Count   int      `json:"count"`
		Entries []result `json:"entries"`
	}
)

const url = "https://api.publicapis.org/entries"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	var r response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	fmt.Println(r)
}