package main

import (
  "fmt"
)

func hello() (string, int) {
  return "Hello, Testing!", 1
}

func main() {
  fmt.Println(hello())
}
