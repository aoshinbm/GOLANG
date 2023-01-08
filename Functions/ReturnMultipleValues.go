package main

import "fmt"

func operations(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub //return multiple values
	//also known as terminating statement
}

func main() {

	sum, diff := operations(100, 54)
	fmt.Println(sum, diff)
}
