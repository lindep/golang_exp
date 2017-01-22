package main

import "fmt"

//!+
// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func firstFun() func() string {
  return func() string {
      return "second func"
    }
}

func main() {
  s := firstFun()
  fmt.Printf("data type = %T\n", s)
  fmt.Println(s())
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
