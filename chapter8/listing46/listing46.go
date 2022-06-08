package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: curl <url> <output-file>")
		os.Exit(1)
	}
}

func main() {
	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	outfile := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error while getting data from url: %v\n", err)
	}

	defer resp.Body.Close()

	outhandle, err := os.Create(outfile)
	if err != nil {
		log.Fatalf("Error while creating outfile: %v\n", err)
	}

	defer outhandle.Close()

	// tee to both Stdout and the outfile
	dst := io.MultiWriter(outhandle, os.Stdout)
	io.Copy(dst, resp.Body)
}