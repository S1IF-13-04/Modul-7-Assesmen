package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&k)

	peti := k / 800

	karung := k % 800 / 100

	ikat := k % 800 % 100 / 10
	sisa := k % 800 % 100 % 10

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisa)
}
