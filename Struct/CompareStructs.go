package main

import "fmt"

type s1 struct {
	x int
}
type s2 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c2 := s2{x: 5}

	if c == c2 {
		fmt.Println("yea Equal")
	}
}

/*
bcuz struct instances are not of the same struct type
i.e one is of type s1 n another is of type s2
not possible to compare struct of different types
attempting will result in compile time error
OUTPUT:-
invalid operation: c == c2 (mismatched types s1 and s2)
*/
