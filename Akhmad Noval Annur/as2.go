package main

import "fmt"

func main() {
	var x, y, totalBakteri int

	fmt.Scan(&x, &y)

	totalBakteri = 1

	for i := x; i <= y; i++ {
		totalBakteri *= i
	}

	fmt.Println(totalBakteri)
	}