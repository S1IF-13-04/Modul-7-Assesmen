package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan bilangan ganjil tersebut: ")
	fmt.Scan(&n)

	fmt.Print("Output: ")
	for i := 1; i <= n*2; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
}
