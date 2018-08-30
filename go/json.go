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

type InviteStatData struct {
    referred        int
    earnedCurrency  string
    earned          int
}

//func main() {
//    p := person{Name: "majun"}
//    b, _ := json.Marshal(p)
//    fmt.Println(b)
//    fmt.Printf("%T\n", b) // []uint8
//    os.Stdout.Write(b)
//}

func main() {
    value := InviteStatData{}
    value.referred = 0
    value.earnedCurrency = "WON"
    value.earned = 0
    fmt.Println(value)

    jsonData, _ := json.Marshal(value)
    fmt.Println(jsonData)
    fmt.Printf("%T\n", jsonData) // []uint8
    os.Stdout.Write(jsonData)
}
