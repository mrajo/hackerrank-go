package main
import "fmt"

func main() {
	var n, a, b, sum int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%v %v", &a, &b)
		sum = a + b
		fmt.Println(sum)
	}
}