package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["apple"] = 5
	map1["orange"] = 6
	map1["avocado"] = 7
	displayMap(map1)

	map2 := map[int]string{1: "joey", 2: "gooey", 100: "booey", 11: ""}
	displayMap(map2)

	value, exists := map2[1]
	if exists {
		fmt.Printf("%v => %v\n", 1, value)
	}

	value, exists = map2[10]
	if exists {
		fmt.Printf("%v => %v\n", 1, value)
	}

	value, exists = map2[11]
	if exists {
		fmt.Printf("%v => %v\n", 11, value)
	}

	delete(map2, 11)
	displayMap(map2)

}

func displayMap[K comparable, V comparable](m map[K]V) {
	for k, v := range m {
		fmt.Printf("%v => %v\n", k, v)
	}
}