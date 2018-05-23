package main

import . "fmt"

type person struct {
    name string
    age int
}

func info(p *person) {
    Println(p.name, p.age)
}

func main() {
    p := &person{
        name: "majun",
        age: 31,
    }

    info(p)

    Println(p)
}
