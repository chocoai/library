package main

import(
    "fmt"
    "sync"
)

//sync.Pool缓存的期限两次gc之间
//不可以使用sync.Pool去实现一个socket连接池
func main() {
    p := &sync.Pool{
        New: func() interface{} {
            return "Hello, Beijing"
        },
    }

    p.Put("Hello, World")
    fmt.Println(p.Get().(string))
    fmt.Println(p.Get().(string))
}
