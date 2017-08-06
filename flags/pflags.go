package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
  user string
  date string
)

func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	for _, u := range users {
		fmt.Println(`Username:	`, u)
		fmt.Println("")
	}

}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
  flag.StringVarP(&date, "date", "d", "", "Date")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
