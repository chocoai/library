package main

import (
    "fmt"
    "reflect"
)

type Foo struct {
    name string
    age int
}

func main() {
    var i int = 12
    var f float32 = 1.2
    var s []string = []string{"a", "b"}
    var foo = Foo{"majun", 32}

    fmt.Println(reflect.TypeOf(i), reflect.TypeOf(i).String())
    fmt.Println(reflect.TypeOf(f), reflect.TypeOf(f).Name())
    fmt.Println(reflect.TypeOf(s), reflect.TypeOf(s).String())
    fmt.Println(reflect.TypeOf(foo), reflect.TypeOf(foo).Name())

    typ := reflect.TypeOf(foo)
    for i := 0; i < typ.NumField(); i++ {
        field := typ.Field(i)
        fmt.Println(field.Name, field.Type)
    }

    field_name, _ := typ.FieldByName("name")
    fmt.Println(field_name.Name, field_name.Type)
    field_age, _ := typ.FieldByName("age")
    fmt.Println(field_age.Name, field_age.Type)
}
