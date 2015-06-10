package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Split(bufio.ScanWords);

	scanner.Scan();
	t, _ := strconv.Atoi(scanner.Text());

	for i := 0; i < t; i++ {
		scanner.Scan();
		n, _ := strconv.Atoi(scanner.Text());

		var max_ending_here int;
		var max_so_far int;
		var positive_sum int;
		var smallest_negative int;

		positive_sum = 0;
		smallest_negative = 0;

		for j := 0; j < n; j++ {
			scanner.Scan();
			a, _ := strconv.Atoi(scanner.Text());

			// contiguous (Kadane algorithm)
			if j == 0 {
				max_ending_here = a;
				max_so_far = a;
			} else {
				if a >= max_ending_here + a {
					max_ending_here = a;
				} else {
					max_ending_here += a;
				}

				if max_so_far < max_ending_here {
					max_so_far = max_ending_here;
				}
			}

			// non-contiguous
			if a > 0 {
				positive_sum += a;
			} else if smallest_negative == 0 || a > smallest_negative {
				smallest_negative = a;
			}
		}

		var non_contiguous int;
		if positive_sum == 0 {
			non_contiguous = smallest_negative;
		} else {
			non_contiguous = positive_sum;
		}

		fmt.Printf("%d %d\n", max_so_far, non_contiguous);
	}
}