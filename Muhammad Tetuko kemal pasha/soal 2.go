package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("masukan x dan y :")
	var total int = 1
	fmt.Scan(&x, &y)
	for i := x; i <= y; i++ {
		total *= i
	}
	fmt.Println(total)
}
