package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	peti := n / 800
	sisa := n % 800
	karung := sisa / 100
	sisa = sisa % 100
	ikat := sisa / 10
	keping := sisa % 10

	fmt.Printf("%d peti, %d karung, %d ikat, %d keping", peti, karung, ikat, keping)
}
