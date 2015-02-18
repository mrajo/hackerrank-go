package main
import "fmt"

func main() {
	var n, loner int;

	fmt.Scan(&n);
	loner = 0;

	for i := 0; i < n; i++ {
		var a int;
		fmt.Scan(&a);
		loner ^= a;
	}

	fmt.Println(loner);
}
