package main

import (
	"fmt"
	"time"
)

func main() {
	//call the method
	//goroutine is created by jz adding the go keyword
	go greeter("hello")
	greeter("worldFolkzzz")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Microsecond) //jz a hack not a ideal method
		fmt.Println(s)
	}
}
