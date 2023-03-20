package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello"
	var ptr *string = &s
	fmt.Println(s)
	*ptr += strings.ToUpper(s)
	fmt.Println(s)
}
