package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		ganjil := 2*i - 1
		fmt.Print(ganjil, " ")
	}
	fmt.Println()
}
