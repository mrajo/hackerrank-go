package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		t int
		nstr string
		nval int64
		count int
	);

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&nstr);
		nval, _ = strconv.ParseInt(nstr, 10, 64);
		count = 0;
		for j := 0; j < len(nstr); j++ {
			digit, _ := strconv.ParseInt(nstr[j:j+1], 10, 0);
			if digit != 0 && nval % digit == 0 {
				count += 1;
			}
		}
		fmt.Println(count);
	}
}