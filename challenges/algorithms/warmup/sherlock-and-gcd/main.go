package main
import "fmt"

func gcd(a int, b int) int {
	for b != 0 {
		t := b;
		b = a % b;
		a = t;
	}

	return a;
}

func main() {
	var t int;

	fmt.Scan(&t);

	for i := 0; i < t; i++ {
		var n, _gcd int;

		fmt.Scan(&n);
		_gcd = 0;

		for i := 0; i < n; i++ {
			var a int;
			fmt.Scan(&a);

			if _gcd == 0 {
				_gcd = a;
			} else {
				_gcd = gcd(_gcd, a);
			}
		}

		if _gcd == 1 {
			fmt.Println("YES");
		} else {
			fmt.Println("NO");
		}
	}
}