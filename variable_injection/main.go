package main

import "fmt"

var (
	Version = "unknown"
	Build   = "unknown"
)

func main() {
	fmt.Printf("Version = %s\nBuild = %s\n", Version, Build)
}
