package main

import "fmt"

type Greeting func(name string) string

// example 1
func say_1(g Greeting, n string) {
    fmt.Println(g(n))
}

// example 2
func (g Greeting) say_2(n string) {
    fmt.Println(g(n))
}

func english(name string) string {
    return "Hello, " + name
}

func main() {
    say_1(english, "World")
    g := Greeting(english)
    g.say_2("World")
}
