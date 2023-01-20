package main

import (
	"fmt"
	"sync"
)

var x int = 0

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	wg.Add(2)

	go inc()
	go inc()
	//go increm()

	wg.Wait()

	fmt.Println(x)
}

func inc() {
	mut.Lock()
	x = x + 10
	mut.Unlock()

	wg.Done()
}

/*
func increm() {
	x = x + 20
	wg.Done()
}
*/
