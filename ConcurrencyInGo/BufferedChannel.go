package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2)
	//producer
	messages <- "channel"
	messages <- "buffered"
	messages <- "12-1-23"

	//consumer
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
