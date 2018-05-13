package logging

import (
	"fmt"
	"os"
	"time"
)

// GetDateTimeNowInString ...
func GetDateTimeNowInString() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

// LOGPREFIXLOG ...
const LOGPREFIXLOG string = "LOG"

// ERRORPREFIXLOG ...
const ERRORPREFIXLOG string = "ERR"

// DEBUGPREFIXLOG ...
const DEBUGPREFIXLOG string = "DEBUG"

// DEFAULTLOGFILEPATH ...
const DEFAULTLOGFILEPATH string = "helloworld-service.log"

var logFileName = DEFAULTLOGFILEPATH

// SetLogFileName ...
func SetLogFileName(filename string) {
	logFileName = filename
}

func Debug(msg string) {
	internalLog(msg, DEBUGPREFIXLOG)
}

func Log(msg string) {
	internalLog(msg, LOGPREFIXLOG)
}

func Error(msg string) {
	internalLog(msg, LOGPREFIXLOG)
}

func internalLog(msg string, prefix string) {
	// prepare the message
	output_msg := fmt.Sprintf("[%s] [%s] %s", prefix, GetDateTimeNowInString(), msg)

	// print to screen and append to log file
	fmt.Println(output_msg)
	appendToLogFile(output_msg, prefix)
}

func appendToLogFile(output_msg string, prefix string) {
	// append log to file
	f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, output_msg)
}
