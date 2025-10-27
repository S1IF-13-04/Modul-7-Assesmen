package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	i := 2
	count := 0
	for count < n {
		fmt.Print(i)
		count++
		if count < n {
			fmt.Print(", ")
		}
		i += 2
	}
	fmt.Println()
}