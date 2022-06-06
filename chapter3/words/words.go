package words

import "strings"

func CountWords(text string) int {
	return len(strings.Split(text, " "))
}