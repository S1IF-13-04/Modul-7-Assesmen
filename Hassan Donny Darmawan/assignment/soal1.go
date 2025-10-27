package main
import "fmt"

func main(){
	var i,n int
	fmt.Print("masukan berapa angka yang ingin dicetak: ")
	fmt.Scan(&n)

	for i=1; i<=n*2; i+=2 {
		fmt.Print(i," ")
	}
}