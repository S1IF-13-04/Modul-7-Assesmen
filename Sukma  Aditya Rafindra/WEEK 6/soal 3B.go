package main

import "fmt"

func main() {
	var keping int
	fmt.Print("masukan angka: ")
	fmt.Scan(&keping)

	peti := keping / 800
	karung := (keping % 800) / 80
	ikat := (keping % 80) / 8
	sisaKeping := keping % 8

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisaKeping)
}
