package main

import "fmt"

func getUtopianTreeHeight(cycles int) int {
	height := 1
	for i := 1; i <= cycles; i++ {
		if i % 2 != 0 {
			height *= 2
		} else {
			height += 1
		}
	}
	return height
}

func main() {
	var n, cycles, height int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&cycles)
		height = getUtopianTreeHeight(cycles)
		fmt.Println(height)
	}
}