package main
import "fmt"
func main() {
	var n, i int

	fmt.Print("masukkan n: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++{
		fmt.Print(i*2, " ")
	}
}