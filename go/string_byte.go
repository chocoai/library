package main

import "fmt"

func main() {
	//var str string = "test"
	//var data []byte = []byte(str)
	//fmt.Println(data)

	var data [10]byte
	data[0] = 'T'
	data[1] = 'E'
	var str string = string(data[:])
	fmt.Println(str)
}
