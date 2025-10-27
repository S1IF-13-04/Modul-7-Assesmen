package main
import "fmt"
func main() {
    var k int
    fmt.Scan(&k)
    p := k/800
    s := k % 800	
    ka := s / 100
    s = s % 100
    i := s / 10
    ks := s % 10
    fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping", p, ka, i, ks)
}