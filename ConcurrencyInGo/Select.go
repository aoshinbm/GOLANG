package main

import "fmt"

func main() {

	number := make(chan int)
	message := make(chan string)

	go channelMessage(message)
	go channelNumber(number)

	select {

	case firstChannel := <-number:
		fmt.Println("Data From Number Channel:", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Data From Message Channel:", secondChannel)
	}
}

// goroutine to send string data to channel
func channelMessage(message chan string) {
	message <- "Learning Go select "
}

// goroutine to send integer data to channel
func channelNumber(number chan int) {
	number <- 4
}
