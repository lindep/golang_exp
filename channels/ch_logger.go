package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel zero memory allocated

func main() {
	go logger()
	logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting donw"}
	time.Sleep(100 * time.Millisecond)
	// shutdown logger
	doneCh <- struct{}{} // empty struct, zero memory allocation
}

func logger() {
	for {
		// this is a blocking select statement because no default case.
		select {
		case entry := <-logCh:
			// Format("2017-12-12T10:10:10")
			fmt.Printf("%v - [%v]%v\n", entry.time, entry.severity, entry.message)
		case <-doneCh:
			break
			// default:
			// 	fmt.Println("Add this to make this a non-blocking Select")
		}
	}
}
