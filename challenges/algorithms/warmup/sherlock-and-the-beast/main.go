package main
import "fmt"
import "strings"

func solve(n int) string {
	if n < 3 {
		return "-1";
	}

	var fives, threes int;

	fives = (n / 3) * 3;
	threes = n - fives;

	for {
		if threes % 5 == 0 && threes <= n {
			if fives < 0 {
				fives = 0;
			}
			return strings.Repeat("5", fives) + strings.Repeat("3", threes);
		}

		if fives < 0 {
			return "-1";
		}

		fives -= 3;
		threes += 3;
	}
}

func main() {
	var t int;

	fmt.Scan(&t);

	for i := 0; i < t; i++ {
		var n  int;

		fmt.Scan(&n);

		fmt.Println(solve(n));
	}
}