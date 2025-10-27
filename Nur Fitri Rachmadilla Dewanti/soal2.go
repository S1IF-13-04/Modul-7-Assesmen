package main
import "fmt"
func main() {
	var x, y, i int
	var total  int

	fmt.Print("masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("masukkan y: ")
	fmt.Scan(&y)

	total = 1
	for i = x; i <=y; i++{
		total = total * i
	}
	fmt.Print("total: ", total)
}