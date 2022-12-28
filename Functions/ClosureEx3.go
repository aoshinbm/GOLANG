package main

import "fmt"

func Filter(intSlice []int, f func(int) bool) []int {

	filtered := make([]int, 0)

	for _, v := range intSlice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
func main() {
	testSlice := []int{3, -1, -5, 2, 4, -8, 70, 20}

	evenNumbers := Filter(testSlice, func(num int) bool {
		flag := false
		if num%2 == 0 {
			flag = true
		}
		return flag
	})
	fmt.Println(evenNumbers)
}
