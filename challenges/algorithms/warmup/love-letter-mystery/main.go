package main
import "fmt"
import "math"

func main() {
	var n, left_pointer, right_pointer int
	var ops float64
	var str string

	fmt.Scan(&n)

	// s := "abcrtz"
	// fmt.Println(reflect.TypeOf(s[0]))
	// fmt.Println(int32(s[2]) - int32(s[0]))
	// fmt.Println(int32(s[0]) - int32(s[2]))
	// fmt.Println(math.Abs(float64(s[0] - s[2])))
	// fmt.Println(math.Abs(float64(int32(s[0]) - int32(s[2]))))

	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		left_pointer = 0
		right_pointer = len(str) - 1
		ops = 0

		for right_pointer > left_pointer {
			ops += math.Abs(float64(int32(str[left_pointer]) - int32(str[right_pointer])))
			right_pointer--
			left_pointer++
		}

		fmt.Println(ops)
	}
}