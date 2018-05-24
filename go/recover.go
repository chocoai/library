package main

import "fmt"

func local_panic() {
    var x = 30
    var y = 0

    var z = x / y
    fmt.Println(z)
}

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    local_panic()

    fmt.Println("...")
}
