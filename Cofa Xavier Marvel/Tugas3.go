package main

import "fmt"

func main() {

	var In int
	fmt.Scan(&In)

	peti := (In / 800)
	karung := (In % 800) / 80
	ikat := ((In % 800) % 80) / 8
	keping := ((In % 800) % 80) % 8

	fmt.Println(peti, "peti", karung, "karung", ikat, "ikat", keping, "keping")

}
