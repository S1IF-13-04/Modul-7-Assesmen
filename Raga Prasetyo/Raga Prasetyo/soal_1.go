package main
import "fmt"

func main() {
	var n int

	fmt.Print("Input: ")
	fmt.Scan(&n)


	fmt.Print("Output: ")

	for i := 1; i <= n; i++ {
		bilanganGanjil := 2*i - 1
		fmt.Print(bilanganGanjil, " ")
	}
	fmt.Println()
}
