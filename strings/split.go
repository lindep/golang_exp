package main

import (
	"bufio"
	_ "encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lineObj := make(map[string]string)

		newJSONStr, delim := "", ""

		fmt.Fprintln(os.Stdout, "---Start line---")
		//fmt.Println(scanner.Text()+ "new line") // Println will add back the final '\n'
		//raw := scanner.Bytes()
		fmt.Fprintln(os.Stdout, scanner.Bytes())
		numBytes := myOut(scanner.Text())
		splitText := strings.Split(scanner.Text(), "|")
		fmt.Fprintln(os.Stdout, "\nNumber of values in list =", len(splitText))
		for i, chunk := range splitText {
			subValObj := make(map[int]string)

			fmt.Println(i, chunk)
			keyValue := strings.Split(chunk, "=")
			fmt.Printf("%d KEY = %s\n", i, keyValue[0])

			if strings.Contains(keyValue[1], ",") {
				subJSONStr, subDelim := "", ""
				for valIdx, value := range strings.Split(keyValue[1], ",") {
					fmt.Printf("\nfound sub keys\n")
					fmt.Printf("%d val = %s\n", valIdx, value)
					subValObj[valIdx] = value
					subJSONStr = subJSONStr + subDelim + strconv.Itoa(valIdx) + ":\"" + value + "\""
					subDelim = ","
				}
				subJSONStr = "[" + subJSONStr + "]"
				fmt.Println(subValObj)
				fmt.Println(subJSONStr)
				newJSONStr = newJSONStr + delim + "\"" + keyValue[0] + "\":" + subJSONStr
				//lineObj[keyValue[0]] = subValObj

				// for n, word := range strings.Split(chunk, "=") {
				// 	fmt.Println(n, word)
				// }
			} else {
				fmt.Printf("%d VALUE = %s\n", i, keyValue[1])
				lineObj[keyValue[0]] = keyValue[1]
				newJSONStr = newJSONStr + delim + "\"" + keyValue[0] + "\":\"" + keyValue[1] + "\""
			}
			delim = ","

		}
		newJSONStr = "{" + newJSONStr + "}"
		fmt.Println(newJSONStr)
		fmt.Println(lineObj)
		fmt.Printf("\nsplitText type = %T\n", splitText)

		fmt.Fprintln(os.Stdout, " Written bytes", numBytes)

		fmt.Printf("%q\n", strings.Split(scanner.Text(), "|"))
		fmt.Printf("%q\n", strings.SplitAfter(scanner.Text(), "|"))

		fmt.Fprintln(os.Stdout, "split", strings.Split(scanner.Text(), "|"))
		fmt.Fprintln(os.Stdout, "splitAfter", strings.SplitAfter(scanner.Text(), "|"))

		fmt.Fprintln(os.Stdout, "\n====End line====")

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
