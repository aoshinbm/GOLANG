package main

import "fmt"

type s1 struct {
	x int
}

func main() {
	c1 := s1{x: 5}
	c2 := s1{x: 5}
	c3 := s1{x: 6}

	if c1 == c2 {
		fmt.Println("c1 n c2 have same values")
	}
	if c2 != c3 {
		fmt.Println("c1 n c2 have different values")
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
