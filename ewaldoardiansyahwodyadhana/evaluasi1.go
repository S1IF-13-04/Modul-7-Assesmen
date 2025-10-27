package main

import "fmt"

func main() {
	var n int

	fmt.Print("Input: ")
	fmt.Scan(&n)

	odd := 1

	fmt.Print("Output: ")

	for i := 0; i < n; i++ {
		fmt.Print(odd)

		if i < n-1 {
			fmt.Print(" ")
		}

		odd += 2
	}

	fmt.Println()
}
