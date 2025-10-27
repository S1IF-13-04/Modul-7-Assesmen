package main
import "fmt"
func main() {
	var n int
	var peti, karung, ikat, keping int

	fmt.Scan(&n)

	peti = n / 800
	karung = (n % 800) / 80
	ikat = n % 800 % 80 / 8
	keping = n % 800 % 80 % 8
	fmt.Println(peti, "peti", karung, "karung", ikat, "ikat", keping, "keping")
}
