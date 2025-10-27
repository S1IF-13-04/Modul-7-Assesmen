package main
import "fmt"
func main() {
	var b, peti, karung, ikat, keping int
	fmt.Scan(&b)
	
	peti = b / 800
	karung = b % 800 / 80
	ikat = b % 800 % 80 / 8
	keping = b % 800 % 80 % 8

	fmt.Println("peti", peti)
	fmt.Println("karung", karung)
	fmt.Println("ikat", ikat)
	fmt.Println("keping", keping)
}