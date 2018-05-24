package main

import (
    "os"
    "fmt"
    "encoding/json"
)

// 成员变量必须大写
type person struct {
    Name string `json:"name"`
    Age int `json:"age,omitempty"`
}

func main() {
    p := person{Name: "majun"}
    b, _ := json.Marshal(p)
    fmt.Println(b)
    fmt.Printf("%T\n", b) // []uint8
    os.Stdout.Write(b)
}
