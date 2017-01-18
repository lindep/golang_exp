package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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
	
	for scanner.Scan() {
		//lineObj := make(map[string]string)

		newJsonStr, delim := "", ""

		//fmt.Fprintln(os.Stdout, "---Start line---")
		splitText := strings.Split(scanner.Text(), "|")
		for _, chunk := range splitText {
			//subValObj := make(map[int]string)
			keyValue := strings.Split(chunk, "=")
			if strings.Contains(keyValue[1], ",") {
				subJsonStr, subDelim := "", ""
				for valIdx, value := range strings.Split(keyValue[1], ",") {
					//subValObj[valIdx] = value
					subJsonStr = subJsonStr+subDelim+"{\""+strconv.Itoa(valIdx)+"\":\""+value+"\"}"
					subDelim = ","
				}
				subJsonStr = "["+subJsonStr+"]"
				newJsonStr = newJsonStr+delim+"\""+keyValue[0]+"\":"+subJsonStr
			} else {
				//lineObj[keyValue[0]] = keyValue[1]
				newJsonStr = newJsonStr+delim+"\""+keyValue[0]+"\":\""+keyValue[1]+"\""
			}
			delim = ","
			
		}
		newJsonStr = "{"+newJsonStr+"}"

		myOut(newJsonStr+"\n")

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}


