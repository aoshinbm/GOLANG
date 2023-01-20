package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("success wrote", i, "to channel")
	}
	close(ch)
}
func main() {

	cha := make(chan int, 2)
	go write(cha)

	time.Sleep(2 * time.Second)
	for v := range cha {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
