package main

import (
	"fmt"
)

func BinarySearch(haystack []int64, needle int64) int64 {
	var lo, hi, mid int64;

	lo = 0;
	hi = int64(len(haystack));

	for lo <= hi {
		mid = (lo + hi) >> 1;

		if haystack[mid] > needle {
			hi = mid - 1;
		} else if haystack[mid] < needle {
			lo = mid + 1;
		} else {
			return mid;
		}
	}

	return -1;
}

func ReadToArray(n int64) []int64 {
	var a []int64;
	var i, t int64;

	for i = 0; i < n; i++ {
		fmt.Scan(&t);
		a = append(a, t);
	}

	return a;
}

func main() {
	var v, n int64;
	fmt.Scan(&v, &n);

	a := ReadToArray(n);

	fmt.Println(BinarySearch(a, v));
}