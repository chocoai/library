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
            return "MAJUN"
        },
    }

    a := p.Get().(string)
    b := p.Get().(string)
    p.Put("1")
    c := p.Get().(string)
    fmt.Println(a, b, c)
}
