package main

import (
	"fmt"
	"math"
)

func main() {
	var n, ops, sum, i uint64;

	fmt.Scan(&n, &ops);

	sum = 0;

	for i = 0; i < ops; i++ {
		var a, b, k uint64;
		fmt.Scan(&a, &b, &k);

		sum += (b - a + 1) * k;
	}

	average := float64(sum) / float64(n);

	fmt.Println(uint64(math.Floor(average)));
}