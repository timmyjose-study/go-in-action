package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	w "words"
)

func main() {
	if len(os.Args[1:]) == 0 {
		data := []string{}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
		text := strings.Join(data, " ")
		fmt.Printf("There are %d words in the text\n", w.CountWords(text))

	} else {
		filename := os.Args[1]

		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("There are %d words in the text\n", w.CountWords(string(content)))
	}
}