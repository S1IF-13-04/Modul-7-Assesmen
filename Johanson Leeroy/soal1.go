package main

import "fmt"

func main() {
	var n, genap, i int

	fmt.Print("Bilangan: ")
	fmt.Scan(&n)

	for i = n; i > 0; i-- {
		genap += 2
		fmt.Print(genap, " ")
	}
}
