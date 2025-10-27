package main

import "fmt"

func main() {
	var keping int
	fmt.Print("input angka: ")
	fmt.Scan(&keping)

	peti := keping / 800
	karung := keping % 800 / 80
	ikat := keping % 800 % 80 / 8
	kepping := keping % 800 % 80 % 8
	fmt.Print(peti, " peti ", karung, " karung ", ikat, " ikat ", kepping, " keping")
}
