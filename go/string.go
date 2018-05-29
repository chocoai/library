package main

import (
    "fmt"
    "unicode"
)

func goodParamFuncName(name string) bool {
    if name == "" {
        return false
    }

    for _, r := range name {
        switch {
        case r == '_':
        case !unicode.IsLetter(r):
            return false
        }
    }

    return true
}

func main() {
    name := "ma_jun"
    fmt.Println(goodParamFuncName(name))
}
