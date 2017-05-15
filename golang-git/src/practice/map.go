package main

import "fmt"

func main() {
    m := make(map[string]string)

    m["one"] = "1"
    m["two"] = "2"
    m["three"] = "3"
    m["four"] = "4"
    m["five"] = "5"

    delete(m, "two")
    delete(m, "forty") // silent behaviour

    if val, ok := m["two"]; ok {
        fmt.Println("Value of m[\"two\"] =", val)
    }

    printMap(&m)

    fmt.Println()

    printMap(&m)

    fmt.Println()

    printMap(&m)

    fmt.Println()

    for k, v := range m {
        fmt.Println(k, ":", v)
    }
    fmt.Println()

    for k, v := range m {
        fmt.Println(k, ":", v)
    }
    fmt.Println()

    for k, v := range m {
        fmt.Println(k, ":", v)
    }
}


func printMap(m *map[string]string) {
    for k, v := range *m {
        fmt.Println(k, ":", v)
    }
}


