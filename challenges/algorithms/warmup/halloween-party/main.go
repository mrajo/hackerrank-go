package main
import "fmt"

func main() {
	var n int;

	fmt.Scan(&n);

	for i := 0; i < n; i++ {
		var k int;
		fmt.Scan(&k);

		v := k / 2;
		h := v + k % 2;

		fmt.Println(v * h);
	}
}