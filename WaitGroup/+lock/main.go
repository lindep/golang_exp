// no concurrency and parallelism

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	fmt.Printf("Threads %v\n", runtime.GOMAXPROCS(-1))
	// runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello(i)
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello(i int) {
	fmt.Printf("Counter %d %v\n", i, counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
