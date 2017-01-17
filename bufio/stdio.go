package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
)

func myOut(msg string) int {
	//f := bufio.NewWriter(os.Stdout)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	//f.Write(buf)
	num ,err := f.WriteString(msg)
	if (err != nil) {
		return -1
	}
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//         advance, token, err = bufio.ScanWords(data, atEOF)
	//         // if err == nil && token != nil {
	//         //         _, err = strconv.ParseInt(string(token), 10, 32)
	//         // }
	//         return
 //  }
  //bufio.ScanWords
	//scanner.Split(split)
	// When no split default to lines
	for scanner.Scan() {
		//fmt.Println(scanner.Text()+ "new line") // Println will add back the final '\n'
		//raw := scanner.Bytes()
		//fmt.Fprintln(os.Stdout, scanner.Bytes())
		numBytes := myOut(scanner.Text())
		fmt.Fprintln(os.Stdout, " Written bytes", numBytes)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}


