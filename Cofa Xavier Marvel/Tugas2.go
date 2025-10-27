package main

import "fmt"

func main() {

	var n, m int
	fmt.Scan(&n, &m)

	result := 1

	for j := n; j <= m; j += 1 {
		result = result * j

	}
	fmt.Println(result)
}
