package main

import "fmt"

func main() {
	num := 108
	func() {
		fmt.Println(num)
	}()
}
