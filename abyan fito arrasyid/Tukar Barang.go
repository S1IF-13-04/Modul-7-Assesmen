package main

import "fmt"

func main() {
	var keping int
	fmt.Scan(&keping)

	peti := keping / 800
	sisa := keping % 800
	karung := sisa / 100
	sisa %= 100
	ikat := sisa / 10
	sisa %= 10

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisa)
}
