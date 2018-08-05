package main

import "fmt"
import "encoding/json"
import (
    "app/model"
    )

func main() {
    model.Initialize()

    s, _ := json.Marshal(model.Rnet)
    fmt.Println(string(s))
}
