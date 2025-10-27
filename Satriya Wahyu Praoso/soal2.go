package main
import "fmt"
func main () {
	var x, y int
	fmt.Scan (&x, &y)
	bakteri := 1
	for i := x; i <= y; i++ {
		bakteri *=i
	}
	fmt.Print(bakteri)
}