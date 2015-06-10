package main

import (
	"fmt"
	"math"
)

func ReadToArray(n int) []float64 {
	var a []float64;
	var i int;
	var t float64;

	for i = 0; i < n; i++ {
		fmt.Scan(&t);
		a = append(a, t);
	}

	return a;
}

func MaxSubarray(a []float64) int64 {
	var max_ending_here, max_so_far float64;
	max_ending_here = a[0];
	max_so_far = a[0];

	for i := 1; i < len(a); i++ {
		max_ending_here = math.Max(a[i], max_ending_here + a[i]);
		max_so_far = math.Max(max_so_far, max_ending_here);
	}

	return int64(max_so_far);
}

// sum of all positive numbers or largest negative
func MaxSubarrayNonContiguous(a []float64) int64 {
	var positive_sum, smallest_negative float64;

	positive_sum = 0;
	smallest_negative = 0;

	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			positive_sum += a[i]
		} else if smallest_negative == 0 || a[i] > smallest_negative {
			smallest_negative = a[i];
		}
	}

	if positive_sum == 0 {
		return int64(smallest_negative);
	}
	return int64(positive_sum);
}

func main() {
	var t, n int;
	fmt.Scan(&t);

	for i := 0; i < t; i++ {
		fmt.Scan(&n);
		a := ReadToArray(n);
		fmt.Print(MaxSubarray(a));
		fmt.Print(" ");
		fmt.Println(MaxSubarrayNonContiguous(a));
	}
}