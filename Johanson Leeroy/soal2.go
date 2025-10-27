package main

import "fmt"

func main() {
	var x, y, bakteri, i int
	fmt.Print("Hari pertama: ")
	fmt.Scan(&x)
	fmt.Print("hari terakhir: ")
	fmt.Scan(&y)
	bakteri = 1

	for i = x; i <= y; i++ {
		bakteri *= i
	}
	fmt.Print("Jumlah bakteri pada hari terkahir: ", bakteri)
}
