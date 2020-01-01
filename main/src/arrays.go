package main

import "fmt"

// array
func printArray(arr [5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func printGrid(grid [][]int) {
	//	for i := 0; i < len(grid[0]); i++ {
	//		for j, v := range grid[i] {
	//			fmt.Println(i, j, v)
	//		}
	//	}

	for i, v := range grid {
		for j, k := range v {
			fmt.Println(i, j, k)
		}
	}
}

func main() {

	// array
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for i := range arr3 {
		fmt.Printf("%d ", arr3[i])
	}

	for _, v := range arr3 {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")

	printArray(arr3)
	printGrid(grid)
}
