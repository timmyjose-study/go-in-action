package main

import "fmt"

func main() {
	slice1 := make([]int, 5) // len = 5, cap = 5
	displaySlice(slice1)

	slice2 := make([]int, 5, 10) // len = 5, cap = 10
	displaySlice(slice2)

	// slice literal
	slice3 := []int{1, 2, 3, 4, 5} // len = 5, cap = 5
	displaySlice(slice3)

	slice4 := []int{99: 0} // len = 100, cap = 100
	displaySlice(slice4)

	// nil slice
	var slice5 []int // len = 0, cap = 0
	displaySlice(slice5)

	// nil slice
	slice6 := []int{}
	displaySlice(slice6)

	// also nil slice
	slice7 := make([]int, 0)
	displaySlice(slice7)

	slice := make([]int, 0)
	displaySlice(slice)
	slice = append(slice, 100)
	slice = append(slice, 200)
	slice = append(slice, 300)
	slice = append(slice, 400)
	slice = append(slice, 500)
	displaySlice(slice)

	subSlice := slice[2:5]
	displaySlice(subSlice)

	subSlice[0] = 999
	displaySlice(slice)

	// a slice is already a reference type, so copying the slice by value
	// is cheap unlike in the case of arrays
	matrix := [][]int{{-1, -2}, {-3, -4}, {-5, -6}}
	displayMatrix(matrix)
}

func displayMatrix(mat [][]int) {
	for _, row := range mat {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}

func displaySlice(slice []int) {
	fmt.Printf("len = %d, capacity = %d\n", len(slice), cap(slice))
	for _, e := range slice {
		fmt.Printf("%d ", e)
	}
	fmt.Println()
}