package main

import "fmt"

func firstFun() func() string {
  return func() string {
      return "second func"
    }
}

func main() {
  dirs := []string{"one","two"}
  var rmdirs []func()
  for _, dir := range dirs {
    fmt.Println(dir)
    dir := dir
    rmdirs = append(rmdirs, func() {
      fmt.Println(dir)
    })
  }

  for _, rmdir := range rmdirs {
    rmdir()
  }

}
