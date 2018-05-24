package main

import "fmt"

func sum(values []int, result_channel chan int) {
    sum := 0
    for _, value := range values {
        sum += value
    }

    result_channel <- sum
}

func main() {
    values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result_channel := make(chan int, 3)

    go sum(values[:len(values)/2], result_channel)
    go sum(values[len(values)/2:], result_channel)
    go sum(values[len(values)/3:], result_channel)

    sum1, sum2, sum3 := <-result_channel, <-result_channel, <-result_channel
    fmt.Println(sum1, sum2, sum3)
}
