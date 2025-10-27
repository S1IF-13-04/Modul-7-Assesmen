package main

import "fmt"

func main() {
	var a, n, j int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&a, &n)
	j = 1

	for i := a; i <= n; i++ {
		j = j * i
	}
	fmt.Print(j)
}
