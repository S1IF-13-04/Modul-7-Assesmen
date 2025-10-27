package main
import "fmt"

func main() {
	var keping int

	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	peti := keping / (8 * 10 * 10)
	sisa := keping % (8 * 10 * 10)

	karung := sisa / (10 * 10)
	sisa = sisa % (10 * 10)

	ikat := sisa / 10
	sisa = sisa % 10

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisa)
}
