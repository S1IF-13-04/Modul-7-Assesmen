package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Input: ")
	fmt.Scan(&x, &y)
	hasil := 1
	for i := x; i <= y; i++ {
		hasil *= i
	}
	fmt.Printf("%d\n", hasil)
}
