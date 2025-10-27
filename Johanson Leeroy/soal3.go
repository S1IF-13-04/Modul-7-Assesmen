package main

import "fmt"

func main() {
	var peti, karung, ikat, keping int
	fmt.Print("Keping: ")
	fmt.Scan(&keping)

	peti = keping / 800
	karung = keping % 800 / 80
	ikat = keping % 800 % 80 / 8
	keping = keping % 800 % 80 % 8

	fmt.Println(peti, "peti,", karung, "karung,", ikat, "ikat dan", keping, "keping")
}
