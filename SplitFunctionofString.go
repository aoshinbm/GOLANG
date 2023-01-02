package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "12,29,65,89"
	var delimiter = ","
	var parts = strings.Split(str, delimiter)
	fmt.Println(parts)
}
