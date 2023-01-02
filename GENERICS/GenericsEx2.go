package main

import "fmt"

func sumNumbers(num1, num2 int) int {
	sum := num1 + num2
	return sum
}
func sum(num1, num2 float64) float64 {
	sum := num1 + num2
	return sum
}

/*
func SumOfNumbers[T any](num1, num2 T) T {
	sum := num1 + num2
	return sum
}
*/
func SumOfTwoNumbers[T int | float64](num1, num2 T) T {
	sum := num1 + num2
	return sum
}

func main() {
	fmt.Println(sumNumbers(3, 7))
	fmt.Println(sum(3.2, 1.7))
	//	fmt.Println(SumOfNumbers(3, 7))
	fmt.Println(SumOfTwoNumbers(3, 7))
	fmt.Println(SumOfTwoNumbers(5.3, 7.6))

}
