package main
import "fmt"

func main() {
	var x int
	var y int
	var hasil int

	fmt.Scan(&x, &y)
	hasil = x * (x + 1) * y

	fmt.Print(hasil)

}