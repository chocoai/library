package main

import (
    "fmt"
)

//指针和非指针类型的receiver的区别：
//指针传递的是对象的引用，非指针传递的是对象的拷贝
type square struct {
    side float64
}

//func (s square) area() float64 {
func (s *square) area() float64 {
    fmt.Printf("%T\n", s)
    return s.side * s.side
}

func main() {
    s := square{10}
    //s := &square{10}
    fmt.Println(s.area())
}
