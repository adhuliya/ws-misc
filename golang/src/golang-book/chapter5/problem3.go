package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        three := i % 3 == 0
        five := i % 5 == 0

        if three && five {
            fmt.Println("FizzBuzz")
        } else if three {
            fmt.Println("Fizz")
        } else if five {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
