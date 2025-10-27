package main

import "fmt"

func main() {
	var tk, p, k, i, keping int
	fmt.Scanln(&tk)

	p = tk / 800
	sisaSetelahPeti := tk % 800

	k = sisaSetelahPeti / 100
	sisaSetelahKarung := sisaSetelahPeti % 100

	i = sisaSetelahKarung / 10
	keping = sisaSetelahKarung % 10

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", p, k, i, keping)
}
