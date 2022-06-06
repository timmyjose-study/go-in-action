package main

import (
	"log"
	_ "matchers" // just to run its init function
	"os"
	"search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Ukraine")
}
