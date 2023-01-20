package main

import (
	"fmt"
	"strings"
)

func modify(s *string) {
	*s = strings.ToUpper(*s)
}
func main() {
	s := "hello"
	fmt.Println(s)
	modify(&s)
	fmt.Println(s)
}
