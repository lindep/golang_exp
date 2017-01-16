
package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	args := os.Args
	fmt.Printf("arg's = %d\n", len(args));
  if len(args) < 1 {
      fmt.Println("Input file is missing.");
      os.Exit(1);
  }
  fileName := os.Args[1]

	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data type = %T\n", data)
	fmt.Printf("read %d bytes: %q\n", count, data[3:count])

	
	fmt.Println("Hello, 世界")
}


