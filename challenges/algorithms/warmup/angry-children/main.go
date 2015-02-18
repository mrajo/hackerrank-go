package main

import (
	"fmt"
	"sort"
)

func spanDistance(a []int) int {
	return a[len(a) - 1] - a[0]
}

func minimumSpanningSlice(a []int, slice_len int) []int {
	sort.Ints(a)
	min_slice := a[:slice_len]
	span_val := spanDistance(min_slice)

	for i := 1; i < len(a) - slice_len - 1; i++ {
		tmp_span := a[i:i + slice_len]
		tmp_val := spanDistance(tmp_span)

		if tmp_val < span_val {
			min_slice = tmp_span
			span_val = tmp_val
		}
	}

	return min_slice
}

func main() {
	var n, kids, size int
	var sizes []int

	fmt.Scanf("%v %v", &n, &kids)

	for i := 0; i < n; i++ {
		fmt.Scan(&size)
		sizes = append(sizes, size)
	}

	min_span_slice := minimumSpanningSlice(sizes, kids)

	fmt.Println(min_span_slice[len(min_span_slice) - 1] - min_span_slice[0])
}
