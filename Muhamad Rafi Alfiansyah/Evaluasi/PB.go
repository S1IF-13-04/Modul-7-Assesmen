package main
import "fmt"
func main() {
    var x, y int
    fmt.Scan(&x, &y)
    h := 1
    for i := x; i <= y; i++ {
        h *= i
    }
	fmt.Print(h)
}