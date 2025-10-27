package main
import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&d)

	a = d / 800
	b = (d % 800) / 100
	c = (d % 100) / 10
	d = d % 10

	fmt.Println(a, "peti", b, "karung", c, "ikat", d, "keping")

}