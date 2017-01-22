package main

import "fmt"

func firstFun(string) func() string {
  return func() string {
      return "second func"
    }
}
func getDir(dir string) func() {
  return func() {
      fmt.Println(dir,"valid variable scope")
    }
}

func main() {
  dirs := [2]string{"one","two"}
  var rmdirsInvalid []func()
  var rmdirsValid []func()
  for _, dir := range dirs {
    fmt.Println("Add dir",dir)
    //dir := dir
    //f := getDir(dir)
    rmdirsInvalid = append(rmdirsInvalid, func() {
      fmt.Println(dir, "invalid variable scope")
    })
    rmdirsValid = append(rmdirsValid, getDir(dir))
  }

  for _, rmdirInvalid := range rmdirsInvalid {
    rmdirInvalid()
  }

  for _, rmdirValid := range rmdirsValid {
    rmdirValid()
  }

}
