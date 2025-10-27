package main

import "fmt"

func main() {

	var peti, karung, ikat, keping int
	fmt.Scan(&keping)

	peti = keping / 800
	karung = (keping % 800) / 100
	ikat = (keping % 800 % 100) / 10
	keping = keping % 800 % 100 % 10

	fmt.Println(peti, "peti,", karung, "karung,", ikat, "ikat", keping, "keping")
}
