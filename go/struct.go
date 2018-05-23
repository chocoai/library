package main

import . "fmt"

type person struct {
    name string
    age int
}

func main() {
    p := &person{
        name: "majun",
        age: 31,
    }

    Println(p)
}
