package main
import "fmt"

func main(){
	var keping int
	fmt.Scan(&keping)
	peti := keping / 800
	karung := (keping % 800) / 100
	ikat := (keping % 100) / 10
	keping = keping % 10
	fmt.Println(peti," peti,", karung," karung,", ikat," ikat, dan", keping," keping")
}