package main

import (
    "fmt"
    "github.com/leonelquinteros/gotext"
)

/* only one language
func main() {
    gotext.Configure("/home/majun/repository/library/go/gettext/locales", "zh_CN", "default")

    fmt.Println(gotext.Get("name"))
}
*/

/* multiple languages
func main() {
    l := gotext.NewLocale("/home/majun/repository/library/go/gettext/locales", "zh_CN")
    l.AddDomain("default")
    l.AddDomain("test")

    fmt.Println(l.Get("name"))
    fmt.Println(l.GetD("test", "name"))
}
*/

func main() {
    l := gotext.NewLocale("/home/majun/repository/library/go/gettext/locales", "zh_CN")
    l.AddDomain("default")
    l.AddDomain("test")

    fmt.Println(l.Get("name"))
    fmt.Println(l.GetD("test", "err_10030"))
}
