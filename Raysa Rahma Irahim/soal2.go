package main

import "fmt"

func main() {
	var awal, akhir int
	
	fmt.Scan(&awal)
	fmt.Scan(&akhir)

	jumlah := 1
	for i := awal; i <= akhir; i++ {
		jumlah *=i
	}
	fmt.Print(" jumlah bakteri terakhir", jumlah)
}