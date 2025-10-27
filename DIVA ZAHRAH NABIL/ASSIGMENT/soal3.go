package main

import "fmt"

func main() {
	var peti, karung, ikat, keping, sisa int

	fmt.Scan(&keping)
	peti = keping / 800
	sisa = keping % 800

	karung = sisa / 100
	sisa = sisa % 100

	ikat = sisa / 10
	sisa = sisa % 10

	fmt.Print(peti, " peti ", karung, " karung ", ikat, " ikat ", sisa, " keping")

}
