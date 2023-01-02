package main

import "fmt"

func reverseLogic(numbers []int) []int {
	//	rev := []int{}
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - 1 - i

		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {

	numbers := []int{1, 2, 4, 5, 6, 7}
	fmt.Println(numbers)
	fmt.Println(reverseLogic(numbers))

}
