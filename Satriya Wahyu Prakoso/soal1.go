package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print(2*i)
		if i < n {
			fmt.Print(", ")
		}
	}
}