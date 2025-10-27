package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 2; i <= n*2; i += 2 {
		fmt.Print(i, " ")

	}
}
