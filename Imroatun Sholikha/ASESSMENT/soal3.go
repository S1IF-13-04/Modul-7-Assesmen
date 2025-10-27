package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	peti := x / 800
	karung := (x % 800) / 80
	ikat := (x % 800) / 100
	keping := x % 8
	fmt.Println(peti, "peti", karung, "karung", ikat, "ikat", keping, "keping")
}
