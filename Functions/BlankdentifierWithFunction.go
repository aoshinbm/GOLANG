package main

import "fmt"

func f() (int, int) {
	return 56, 89
}
func main() {
	a, b := f()
	fmt.Println(a, b)

	/*
		v := f()
		fmt.Println(v)
		assignment mismatch: 1 variable but f returns 2 values
	*/

	//here blank identifier is used bcuz one variable was required
	//i.e to eliminate the other one
	x, _ := f()
	fmt.Println(x)

}
