package main

import (
	"fmt"
)

func hello(chnll chan bool) {
	fmt.Println(" Hello, I am in go routine ")
	chnll <- true
}

func main() {
	chnel := make(chan bool)
	go hello(chnel)
	<-chnel
	fmt.Println("MAin functionm")
}
