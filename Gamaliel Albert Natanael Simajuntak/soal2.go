package main
import "fmt"
 func main() {
	var x, y int
	fmt.Scan(&x,&y)
	for i:=2; i>=1; i-- {
		fmt.Print(x*i, " ", y*i, " ")

	}
}