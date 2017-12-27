package main

import (
  "fmt"
  "github.com/lindep/golang_exp/example/package/hello"
)

func main() {
  fmt.Println("Loading custom package")
  hello.PrintHello()
}
