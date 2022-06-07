package main

import "fmt"

func main() {
	var arr1 [5]int // zeroed
	displayArray(arr1)

	arr2 := [...]int{1, 2, 3, 4, 5}
	displayArray(arr2)

	arr3 := [...]int{1: 100, 4: 200}
	displayArray(arr3)
	arr3[0] = -100
	displayArray(arr3)

	// array of pointers
	arr4 := [5]*int{0: new(int), 1: new(int)}
	*arr4[0] = 100
	*arr4[1] = 200
	displayPtrArray(arr4)

	// assigning arrays
	var sarr1 [5]string
	sarr2 := [...]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	sarr1 = sarr2 // copy

	for i := 0; i < len(sarr1); i++ {
		fmt.Printf("%s, %s\n", sarr1[i], sarr2[i])
	}

	// assigning pinter arrays - simply copies the pointer,
	// not the value pointed to
	var parr1 [3]*string

	parr2 := [...]*string{new(string), new(string), new(string)}
	*parr2[0] = "Red"
	*parr2[1] = "Blue"
	*parr2[2] = "Green"

	parr1 = parr2
	for i := 0; i < len(parr1); i++ {
		if parr1[i] != nil {
			fmt.Printf("%s at %p, %s at %p\n", *parr1[i], parr1[i], *parr2[i], parr2[i])
		} else {
			fmt.Println("nil")
		}
	}

	multiDimArrays()
	passingArraysToFunction()

}

func passingArraysToFunction() {
	// in Go, array are passed by value - the whole array, not the reference
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	processArray(&arr)
}

func processArray(arr *[10]int) {
	for _, n := range arr {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func multiDimArrays() {
	matrix := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

	anotherMatrix := [3][2]int{1: {0: 20, 1: 30}, 2: {1: 100}}
	for _, row := range anotherMatrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

}

func displayPtrArray(arr [5]*int) {
	for _, p := range arr {
		if p == nil {
			fmt.Printf("nil ")
		} else {
			fmt.Printf("%d ", *p)
		}
	}
	fmt.Println()

}

func displayArray(arr [5]int) {
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
	fmt.Println()
}