package main

import (
	"fmt"
	"math"
)

func IsPerfectSquare(n float64) bool {
	return math.Mod(math.Sqrt(n), 1) == 0;
}

func IsFibonacci(n float64) string {
	// 5x^2
	var fxs float64;
	fxs = 5 * math.Pow(n, 2);

	if IsPerfectSquare(fxs + 4) || IsPerfectSquare(fxs - 4) {
		return "IsFibo";
	}
	return "IsNotFibo";
}

func main() {
	var (
		t int
		n float64
	);

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n);
		fmt.Println(IsFibonacci(n));
	}
}