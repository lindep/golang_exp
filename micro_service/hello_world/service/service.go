package service

import "fmt"

// HelloWorld concat name passed and return string
func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}
