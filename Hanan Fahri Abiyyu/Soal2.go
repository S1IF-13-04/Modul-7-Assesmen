package main
import "fmt"
func main() {
	var x, y int
	var total int
	fmt.Scan(&x, &y)
	for i := x; i <= y; i++{
		total = x * (x+1) * y
	}
	fmt.Printf("%d", total )
}
