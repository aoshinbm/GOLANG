package main

import (
	"fmt"
)

func main() {
	var intChnl chan int
	fmt.Printf("Type of intChnl: %T \n", intChnl)

	strngChnl := make(chan string)
	fmt.Printf("Type of strngChnl: %T \n", strngChnl)
}
