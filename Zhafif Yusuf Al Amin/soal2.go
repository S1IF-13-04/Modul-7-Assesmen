package main

import "fmt"

func main() {
	var n, m, hasil int
	fmt.Print("Input: ")
	fmt.Scan(&m,&n)
	hasil = 1
	for iterasi := m; iterasi <= n; iterasi = iterasi + 1 {
		hasil = hasil * iterasi
		
	}
	fmt.Println(hasil)
}
