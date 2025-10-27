package main

import "fmt"

func main() {
	var awal, akhir int
	fmt.Print("Hari awal: ")
	fmt.Scan(&awal)
	fmt.Print("Hari akhir: ")
	fmt.Scan(&akhir)

	jumlah := 1
	for i := awal; i <= akhir; i++ {
		jumlah *=i
	}
	fmt.Print("Jumlah Bakteri Terakhir: ", jumlah)
}
