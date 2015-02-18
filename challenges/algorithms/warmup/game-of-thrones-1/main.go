package main

import "fmt"

func canBePalindrome(letters map[string]int) bool {
	oddCounts := 0

	for _, val := range letters {
		if val % 2 != 0 {
			oddCounts++
		}
		if oddCounts > 1 {
			break
		}
	}

	return oddCounts <= 1
}

func main() {
	var str string
	fmt.Scan(&str)

	counts := map[string]int {}

	for i := 0; i < len(str); i++ {
		counts[string(str[i])]++
	}

	if canBePalindrome(counts) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
