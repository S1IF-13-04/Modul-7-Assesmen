package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	hasil := 1
	for i := n; i <= x; i++ {
		hasil *= i

	}
	fmt.Println(hasil)
}
