package main

import "fmt"

func main() {
    var x [5]int

    fmt.Println(x)

    for i := 0; i < len(x); i++ {
        fmt.Println(x[i])
    }

    fmt.Println()

    for i, val := range x {
        fmt.Println(i, "=", val)
    }
}
