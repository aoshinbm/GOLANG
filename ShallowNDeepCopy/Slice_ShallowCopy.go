//Write a program to demonstrate shallow copy of a slice in Go.

package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	copied := original
	copied[0] = 10
	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)
}
