package main

import "fmt"

func main() {
	var totalkeping, peti, karung, ikat, sisakeping int
	fmt.Scan(&totalkeping)
	peti = totalkeping / 800
	karung = (totalkeping % 800) / 80
	ikat = (totalkeping % 80) / 8
	sisakeping = totalkeping % 8
	fmt.Println(peti, "Peti,", karung, "karung,", ikat, "ikat, dan", sisakeping, "keping")
}
