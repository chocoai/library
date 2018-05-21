package main

import "fmt"

// example 1
func f1() (result int) {
    defer func() {
        result++
    }()
    return 0
}

// example 2
func f2() (result int) {
    t := 5
    defer func() {
        t = t + 5
    }()
    return t
}

// example 3
func f3() (result int) {
    defer func(result int) {
        result = result + 5
    }(result)

    return 1
}

func main() {
    fmt.Println(f1())
    fmt.Println(f2())
    fmt.Println(f3())
}
