package main

import "fmt"

func main() {
	var n, i, hasil int
	fmt.Scan(&n)
	for i = 2; i < n; i += 2 {
		hasil = i
		fmt.Print(hasil)
	}
}
