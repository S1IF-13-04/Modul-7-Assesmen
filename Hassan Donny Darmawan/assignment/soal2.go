package main
import "fmt"

func main (){
	var i,n,hasil,x int
	fmt.Print("masukan hari pertama dan hari terakhir: ")
	fmt.Scan(&x,&n)
	hasil = 1

	for i=x; i<=n; i++ {
		hasil *= i
	}
	fmt.Print(hasil)
}