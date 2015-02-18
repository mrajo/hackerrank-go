package main
import "fmt"

func main() {
	var n int;

	fmt.Scan(&n);

	for i := 0; i < n; i++ {
		var steps, a, b int;
		fmt.Scan(&steps, &a, &b);

		val := (steps - 1) * a;
		fmt.Printf("%d ", val);

		for i := 0; i < steps - 1; i++ {
			val = val - a + b;
			fmt.Printf("%d ", val);
		}

		fmt.Println();
	}
}