package main

import "fmt"

func eventFilter(numbers []int) []int {

	evenNumbers := []int{}
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return evenNumbers
}

func oddFilter(numbers []int) []int {

	oddNumbers := []int{}
	for _, number := range numbers {
		if number%2 != 0 {
			oddNumbers = append(oddNumbers, number)
		}
	}
	return oddNumbers
}

func main() {

	numbers := []int{1, 2, 4, 5, 6, 7, 8, 9, 1, 11, 12, 14, 5, 13, 15, 16, 16, 18, 19}
	evenNumbers := eventFilter(numbers)
	fmt.Println(evenNumbers)

	fmt.Println("---------------------------------------")

	oddNumbers := oddFilter(numbers)
	fmt.Println(oddNumbers)

}

/* function  -- ()

func functionName (paramtere list -- Any data type(User defiend or golang provide) ) retun type {

}

Higher order function --- () --- Abstracting the logic

func functionName (paramtere list -- another func or func's) ) retun func {

}



*/
