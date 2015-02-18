package main

import (
	"fmt"
	"math"
)

func main() {
	var n int;

	fmt.Scan(&n);

	for i := 0; i < n; i++ {
		var n, c, m int;
		fmt.Scan(&n, &c, &m);

		bought := math.Floor(float64(n / c));
		eaten := math.Floor(bought + (bought - 1) / (float64(m) - 1));
		fmt.Println(eaten);
	}
}