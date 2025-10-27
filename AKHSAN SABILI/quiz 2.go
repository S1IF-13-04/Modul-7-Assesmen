package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&a, &n)
	hasil := 1
	for i := a; i <= n; i++ {
		hasil = hasil * i
	}
	fmt.Println(hasil)
}
