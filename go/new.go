package main

import "fmt"

type person struct {
    name string
}

func main() {
    p1 := new(person)
    fmt.Println(p1)

    p2 := person{}
    fmt.Println(p2)
}
