package main

import (
    "fmt"
    "reflect"
    "runtime"
)

type Foo struct {
    name string
    age int
}

type Handler func()

func handler() {
    fmt.Println("handler")
}

func HandlerName(h Handler) string {
    pc := reflect.ValueOf(h).Pointer()
    return runtime.FuncForPC(pc).Name()
}

func main() {
    var i int = 12
    var f float32 = 1.2
    var s []string = []string{"a", "b"}
    var foo = Foo{"majun", 32}

    fmt.Println(reflect.ValueOf(i))
    fmt.Println(reflect.ValueOf(f))
    fmt.Println(reflect.ValueOf(s))
    fmt.Println(reflect.ValueOf(foo))
    fmt.Println(HandlerName(handler))
}
