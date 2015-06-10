package main

import (
	"fmt"
	"strings"
	"strconv"
)

func ReadToArray(n int64) []int64 {
	var a []int64;
	var i, t int64;

	for i = 0; i < n; i++ {
		fmt.Scan(&t);
		a = append(a, t);
	}

	return a;
}

func PrintArray(a []int64) {
	output := "";
	for i := 0; i < len(a); i++ {
		output += strconv.FormatInt(a[i], 10) + " ";
	}

	fmt.Println(strings.TrimSpace(output));
}

func InsertionSort(a []int64) []int64 {
	n := len(a) - 1;
	v := a[n];

	for i := n; i >= 0; i-- {
		if i != 0 && a[i - 1] > v {
			a[i] = a[i - 1];
		} else {
			a[i] = v;
			PrintArray(a);
			break;
		}
		PrintArray(a);
	}
	return a
}

func main() {
	var size int64;
	fmt.Scan(&size);

	a := ReadToArray(size);

	InsertionSort(a);
}