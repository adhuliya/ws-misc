package main

import "fmt"

func main() {
    var a = "anshuman"

    const b = a

    a = "dhuliya"

    fmt.Println(b, a)
}
