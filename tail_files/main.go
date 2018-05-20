package main

import (
	"fmt"
	"log"

	"github.com/hpcloud/tail"
)

func main() {
	t, err := tail.TailFile("/tmp/testtail.log", tail.Config{Follow: true})
	if err != nil {
		log.Fatal(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
