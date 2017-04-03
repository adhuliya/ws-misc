package main

import "fmt"

func main() {
    var f float64
    fmt.Print("Fahrenheit : ")
    fmt.Scanf("%f", &f)
    fmt.Println("Celcius =", (f-32)*5/9)
}
