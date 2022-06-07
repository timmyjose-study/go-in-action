package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./listing34 <url>")
		os.Exit(1)
	}
}

func main() {
	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while getting data from url: %v\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)

	if err = resp.Body.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

}