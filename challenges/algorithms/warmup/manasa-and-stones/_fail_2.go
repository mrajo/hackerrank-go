package main
import "fmt"
import "math"

func main() {
	var n int;

	fmt.Scan(&n);

	for i := 0; i < n; i++ {
		var steps, a, b int;
		fmt.Scan(&steps, &a, &b);

		if a > b {
			a, b = b, a;
		}

		val := (steps - 1) * int(math.Min(float64(a), float64(b)));
		fmt.Printf("%d", val);

		for i := 0; i < steps - 1; i++ {
			val = val + int(math.Abs(float64(a - b)));
			fmt.Printf(" %d", val);
		}

		fmt.Println();
	}
}