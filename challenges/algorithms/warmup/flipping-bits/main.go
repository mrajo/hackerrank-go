package main
import "fmt"

func main() {
    var n int
    var c uint32

    fmt.Scan(&n)

    for i := 0; i < n; i++ {
    	fmt.Scan(&c)
    	fmt.Println(^c)
    }
}