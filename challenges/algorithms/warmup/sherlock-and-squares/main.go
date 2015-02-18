package main
import "fmt"
import "math"

func main() {
	var n int;

	fmt.Scan(&n);

	for i := 0; i < n; i++ {
		var a, b float64;
		fmt.Scan(&a, &b);
		fmt.Println(math.Floor(math.Sqrt(b)) - math.Ceil(math.Sqrt(a)) + 1);
	}
}