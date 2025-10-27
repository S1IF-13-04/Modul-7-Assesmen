package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Input (x y): ")
	fmt.Scan(&x, &y)

	var r int64 = 1

	for i := x; i <= y; i++ {
		r *= int64(i)
	}

	fmt.Printf("Output: %d\n", r)
}
