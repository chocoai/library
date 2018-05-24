package main

import "fmt"

func main() {
    var country_capital map[string]string
    if country_capital == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }

    country_capital = make(map[string]string)

    country_capital["China"] = "北京"
    country_capital["America"] = "华盛顿"

    for country, capital := range country_capital {
        fmt.Println(country, capital)
    }
}
