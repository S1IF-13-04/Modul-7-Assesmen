package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    fmt.Printf("logika\n")
    for i := x; i <= y; i++ {
        if i > x {
            fmt.Print("x")
        }
        fmt.Print(i)
    }
    fmt.Println()

    fmt.Printf("output\n")
    prod := 1
    for i := x; i <= y; i++ {
        prod *= i
    }
    fmt.Println(prod)
}