package main

import (
    "fmt"
    "math"
)

type shape interface {
    area() float64
}

type square struct {
    side float64
}

// 如果receiver是指针类型，传入的必须是指针类型；
// 否则，可以隐式转换，即指针类型自动转换为非指针类型
func (s *square) area() float64 {
    return s.side * s.side
}

type circle struct {
    radius float64
}

func (c *circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func info(z shape) {
    fmt.Println(z)
    fmt.Println(z.area())
}

func totalArea(shapes ...shape) float64 {
    var total_area float64
    for _, s := range shapes {
        total_area += s.area()
    }

    return total_area
}

func main() {
    s := &square{10}
    c := &circle{5}
    info(s)
    info(c)
    totalArea(s, c)
}
