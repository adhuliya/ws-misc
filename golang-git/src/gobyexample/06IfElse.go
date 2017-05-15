package main

import "fmt"

func main() {

    if 7 % 2 == 0 {
        fmt.Println("7 is Even")
    } else {
        fmt.Println("7 is Odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisble by 4")
    }

    if num := 9; num < 0 {
        fmt.Println("num is negetive")
    } else if num = 10; num < 10 {
        fmt.Println("num is a single digit number")
    } else {
        fmt.Println("num is", num)
    }

    x := 10
    x = 20

    fmt.Println(x)

}
