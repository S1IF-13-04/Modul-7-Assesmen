package main
import "fmt"
func main () {
	var a int
	fmt.Scan(&a)
	peti := a / 800
	karung := a % 800 % 10
	ikat := a % 800 / 10 / 10
	keping := a % 800 / 10 % 10
	fmt.Printf("%d peti, %d karung, %d ikat, %d keping", peti, karung, ikat, keping)
}