package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan Input :")
	fmt.Scan(&n)
	for i := 2; (i <= 10); i += 2{
		fmt.Print(i, " ")
	}
}
