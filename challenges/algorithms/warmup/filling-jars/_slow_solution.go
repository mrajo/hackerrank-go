package main

import (
	"fmt"
	"math"
)

/*
This approach won't work because of timeouts with larger data sets
Redo this with a running calculation instead of storing jars in array
*/

func ArraySum(a []uint64) uint64 {
	var sum uint64
	sum = 0;

	for i := 0; i < len(a); i++ {
		sum += a[i];
	}

	return sum;
}

func ArrayAverage(a []uint64) float64 {
	sum := ArraySum(a);
	size := len(a);
	return float64(sum) / float64(size);
}

func main() {
	var n, ops, i uint64

	fmt.Scan(&n, &ops);

	jars := make([]uint64, n);

	for i = 0; i < ops; i++ {
		var a, b, k uint64
		fmt.Scan(&a, &b, &k);

		for j := a; j <= b; j++ {
			jars[j-1] += k;
		}
	}

	fmt.Println(uint64(math.Floor(ArrayAverage(jars))));
}