package main 
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	h := 0
	for i:=1;i<=n;i++{
		h += 2
		fmt.Println(h)
	}
}