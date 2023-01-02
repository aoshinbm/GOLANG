package main

import (
	"fmt"
	"strconv"
)

// first import string conversion package

func main() {
	//convert integer into string
	fmt.Println("Converting int into string :")
	var num int = 25
	var strr string = strconv.Itoa(num)
	fmt.Printf("%q", strr)
}
