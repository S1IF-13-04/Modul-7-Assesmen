package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan Angka: ")
	fmt.Scan(&n)

	for i := 1; i <= n*2; i += 2 {
		fmt.Print(i, " ")
	}
}
