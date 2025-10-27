package main

import "fmt"

func main() {
	var totalQirat int
	fmt.Scan(&totalQirat)

	dinar := totalQirat / 600
	sisa := totalQirat % 600 

	dirham := sisa / 60
	sisa = sisa % 60

	fals := sisa / 6
	qirat := sisa % 6
	fmt.Println(dinar, dirham, fals, qirat)
}