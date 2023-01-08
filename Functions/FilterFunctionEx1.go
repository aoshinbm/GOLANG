package main

import (
	"fmt"
)

func filter(numbers []int, functionName func(int) bool) []int {
	filteredNumbers := []int{}
	for _, number := range numbers {
		if functionName(number) {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return filteredNumbers
}

func evenLogic(number int) bool {
	return number%2 == 0
}

func oddLogic(number int) bool {
	return number%2 != 0
}

func main() {

	numbers := []int{1, 2, 4, 5, 6, 7, 8, 9, 1, 11, 12, 14, 5, 13, 15, 16, 16, 18, 19}

	// Printing the even no from above slice
	evenNumbers := filter(numbers, evenLogic)
	fmt.Println(evenNumbers)

	// Prinitng the odd no form above slice
	oddNumbers := filter(numbers, func(number int) bool {
		return number%2 != 0
	})
	fmt.Println(oddNumbers)

	/* function  -- ()

	   func functionName (paramtere list -- Any data type(User defiend or golang provide) ) retun type {

	   }

	   Higher order function --- () --- Abstracting the logic

	   func functionName (paramtere list -- another func or func's) ) retun func {

	   }



	*/
}
