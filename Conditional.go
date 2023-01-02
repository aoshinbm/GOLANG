package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//generate random values between -5 and 4
	rand.Seed(time.Now().UnixNano())

	num := -5 + rand.Intn(10)

	//With the help of the conditionals, we print a message for all three options
	if num > 0 {
		fmt.Println(num, "The number is positive")
	} else if num == 0 {
		fmt.Println(num, "The number is zero")
	} else {
		fmt.Println(num, "The number is negative")
	}
}
