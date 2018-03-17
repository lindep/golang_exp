
package main

import "fmt"

var (
    a = c - 2
    b = 2
)

func hello() string {
        return "From sandbox!"
}

func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(hello())
}
