package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("listener: Started: listening on :4000")
	http.ListenAndServe("localhost:4000", nil)
}
