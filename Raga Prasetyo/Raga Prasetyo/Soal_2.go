package main
import "fmt"

func main(){
	var x, y int

	fmt.Print("Masukkan (hari awal dan hari akhir): ")
	fmt.Scan(&x, &y)

	hasil := 1
	fmt.Print("Logika: ")

	for i := x; i <= y; i++{
		hasil *= i
		fmt.Print(i)
		if i < y {
			fmt.Print("x")
		}
	}

	fmt.Println()
	fmt.Println("Output:", hasil)
}
