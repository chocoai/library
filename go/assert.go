package main

import "fmt"

func main() {
	var a interface{} = 1
	t, ok := a.(int)
	fmt.Println(t, ok)
}
