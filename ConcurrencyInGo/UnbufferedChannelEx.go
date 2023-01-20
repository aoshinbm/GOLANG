package main

import (
	"fmt"
)

func hello(chnll chan string) {
	fmt.Println(" Hello, I am in go routine ")
	chnll <- " "
}

func main() {
	chnel := make(chan string)
	go hello(chnel)
	<-chnel
	fmt.Println("Main function")
}
