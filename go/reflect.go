package main

import (
    "fmt"
    "reflect"
)

func main() {
    var i int = 12
    var f float32 = 1.2
    var s []string = []string{"a", "b"}

    fmt.Println(reflect.TypeOf(i))
    fmt.Println(reflect.TypeOf(f))
    fmt.Println(reflect.TypeOf(s))
}
