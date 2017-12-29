package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
	"strings"
	"text/scanner"
)

func myOut(msg string) int {
	//f := bufio.NewWriter(os.Stdout)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	//f.Write(buf)
	num, err := f.WriteString(msg)
	if err != nil {
		return -1
	}
	return num
}

func main() {
	var s scanner.Scanner
	s.Filename = "os.Stdin"

	bufScanner := bufio.NewScanner(os.Stdin)
	// split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//         advance, token, err = bufio.ScanWords(data, atEOF)
	//         // if err == nil && token != nil {
	//         //         _, err = strconv.ParseInt(string(token), 10, 32)
	//         // }
	//         return
	//  }
	onPipe := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '|' {
				return i + 1, data[:i], nil
			}
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	//bufio.ScanWords
	bufScanner.Split(onPipe)
	// When no Split() default to lines
	for bufScanner.Scan() {
		//fmt.Println(scanner.Text()+ "new line") // Println will add back the final '\n'
		//raw := scanner.Bytes()
		//fmt.Fprintln(os.Stdout, scanner.Bytes())

		// strings.NewReader(scanner.Text()) -> io.Reader
		s.Init(strings.NewReader(bufScanner.Text()))
		var tok rune
		for tok != scanner.EOF {
			tok = s.Scan()
			fmt.Println("At position", s.Pos(), ":", s.TokenText())
		}

		numBytes := myOut(bufScanner.Text())
		fmt.Fprintln(os.Stdout, " Written bytes", numBytes)
	}
	if err := bufScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
