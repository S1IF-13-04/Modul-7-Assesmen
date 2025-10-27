package main

import (
	"fmt"
)

const (
	K_BALI   = 10
	K_KARUNG = 100
	K_PETI   = 500
)

func main() {
	var totalKeping int
	var peti, karung, bali, keping int
	var sisa int

	fmt.Print("Masukkan total keping: ")
	fmt.Scan(&totalKeping)

	sisa = totalKeping

	peti = sisa / K_PETI
	sisa = sisa % K_PETI

	karung = sisa / K_KARUNG
	sisa = sisa % K_KARUNG

	bali = sisa / K_BALI
	sisa = sisa % K_BALI

	keping = sisa

	fmt.Printf("%d peti, %d karung, %d bali, dan %d keping\n", peti, karung, bali, keping)
}
