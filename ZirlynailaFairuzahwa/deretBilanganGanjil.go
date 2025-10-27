package main
import "fmt"

func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	for i := 1; i <= bilangan; i++ {
		deretGanjil := 2 * i - 1
		fmt.Print(deretGanjil, " ")
	}
}