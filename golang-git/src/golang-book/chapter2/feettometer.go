package main

import "fmt"

func main() {
    var feet float64

    fmt.Print("Feet : ")
    fmt.Scanf("%f", &feet)
    fmt.Println("Meter :", feet * 0.3048)
}
