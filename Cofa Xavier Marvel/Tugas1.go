package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	for j := 2; j <= n*2; j += 2 {
		fmt.Print(j, " ")
	}
}
