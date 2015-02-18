package main

import (
	"fmt"
	"regexp"
)

func main() {
	var n int
	var str, collapse string

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		var reA = regexp.MustCompile(`A+`)
		var reB = regexp.MustCompile(`B+`)
		collapse = reA.ReplaceAllLiteralString(str, "A")
		collapse = reB.ReplaceAllLiteralString(collapse, "B")
		fmt.Println(len(str) - len(collapse))
	}
}