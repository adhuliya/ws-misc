package main

import (
    f "fmt"
)

// this is an enum equivalent in golang
type color int
const (
    RED     = iota
    BLUE
    GREEN
)

func main() {
    f.Println("Start", RED, GREEN, BLUE)
}
