package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	testSlice := []int{10, 10, 30, 30}

	sum := adder()
	for i, val := range testSlice {
		fmt.Println(i, val, sum(val))
	}

}
