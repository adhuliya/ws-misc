package main

import "fmt"

func main() {

    var a string = "Hello"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int // automatically initialized to zero
    fmt.Println(e)

    f := "short"
    fmt.Println(f)
}
