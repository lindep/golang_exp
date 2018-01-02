package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func myOut(msg string) int {
	num, err := io.WriteString(os.Stdout, msg)
	if err != nil {
		return -1
	}
	return num
	/*
		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()
		//f.Write(buf)
		num ,err := f.WriteString(msg)
		if (err != nil) {
			return -1
		}
		return num*/
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if len(scanner.Bytes()) < 1 {
			continue
		}
		newJSONStr, delim := "", ""
		splitText := strings.Split(scanner.Text(), "|")
		for _, chunk := range splitText {
			keyValue := strings.Split(chunk, "=")
			if len(keyValue) != 2 {
				continue
			}

			if strings.Contains(keyValue[1], ",") {
				subJSONStr, subDelim := "", ""
				for valIdx, value := range strings.Split(keyValue[1], ",") {
					//subValObj[valIdx] = value
					subJSONStr = subJSONStr + subDelim + "{\"" + strconv.Itoa(valIdx) + "\":\"" + value + "\"}"
					subDelim = ","
				}
				subJSONStr = "[" + subJSONStr + "]"
				newJSONStr = newJSONStr + delim + "\"" + keyValue[0] + "\":" + subJSONStr
			} else {
				newJSONStr = newJSONStr + delim + "\"" + keyValue[0] + "\":\"" + keyValue[1] + "\""
			}
			delim = ","

		}
		newJSONStr = "{" + newJSONStr + "}"

		myOut(newJSONStr + "\n")

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
