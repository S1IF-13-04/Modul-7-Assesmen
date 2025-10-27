package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah bilangan genap: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d, ", 2*i)
	}

	fmt.Printf("\b\b ")
}
