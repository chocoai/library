package main

import "fmt"

func main() {
    // example 1
    const (
        x = iota
        y
        z
    )

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)

    // example 2
    const (
        a = iota
        b
        c
        d = "ma"
        e
        f = iota
        g
    )

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
    fmt.Println(g)
}
