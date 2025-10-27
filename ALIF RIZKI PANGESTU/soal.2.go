package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan hari x dan y: ")
	fmt.Scan(&x, &y)

	total := 1
	for i := x; i <= y; i++ {
		total *= i
	}
	fmt.Println("Jumlah bakteri pada hari ke-", y, "adalah:", total)
}
