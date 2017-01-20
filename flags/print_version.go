package main

import (
    "flag"
    "fmt"
    "os"
)

const AppVersion = "1.0.0 beta"

func main() {
    version := flag.Bool("v", false, "prints current roxy version")
    flag.Parse()
    if *version {
      fmt.Println(AppVersion)
      os.Exit(0)
    }
    fmt.Println("Hello from main()")
}