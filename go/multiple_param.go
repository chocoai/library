package main

import "fmt"

func main() {
    slice_1 := append([]int{1, 2, 3}, 4, 5, 6)
    slice_2 := append([]int{1, 2, 3}, []int{4, 5, 6}...)
    fmt.Println(slice_1)
    fmt.Println(slice_2)
}
