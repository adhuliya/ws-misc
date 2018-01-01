package main

import "fmt"

func main() {
    var s []int
    s = make([]int, 0, 10)

    fmt.Println("len(s) =", len(s))
    fmt.Println("cap(s) =", cap(s))

    p := append(s, 1)
    _ = append(s, 2)
    _ = append(s, 3)

    s = append(s,4)
    fmt.Println(p)
    fmt.Println(s)

    fmt.Println("len(s) =", len(s))
    fmt.Println("cap(s) =", cap(s))

}
