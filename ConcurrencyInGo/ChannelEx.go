package main

import (
	"fmt"
)

// channel arguemnt shud be passed to go routine
func prod(num1, num2 int, channel chan int) {
	channel <- num1 * num2
}
func main() {
	chnel := make(chan int)
	go prod(10, 20, chnel)
	go prod(200, 800, chnel)
	num1 := <-chnel
	num2 := <-chnel

	fmt.Println(num1, num2, num1*num2)
}
