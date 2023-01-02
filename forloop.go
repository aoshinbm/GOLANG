package main

import "fmt"

func main() {

	sum := 0

	//This is the classic C-style for statement
	//The for statement consists of three parts: the initialization, the condition, and the increment.
	//The initialization part is executed only once.
	//The body of the for statement is executed when the condition is true.
	//If the condition returns false, the for loop is terminated.
	//After the statements in the block are executed,
	//the for loop switches to the third part,
	//where the counter is incremented.
	//The cycle continues until the condition is not true anymore.
	//Note that is it possible to create endless loops.
	for i := 0; i < 10; i++ {
		//sum += i
		sum = sum + i
		fmt.Println(sum)
	}

	fmt.Println("Total sum : ", sum)

	fmt.Println("------------------------------------------------")

	//calculates the sum of array values

	//define an array of values
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	Sum := 0

	//We iterate over the array with the range clause.
	//The range returns the index and the value in each iteration.
	//Since we do not use the index, we specify the discard _ operator.
	//(The Golang documentation calls it the blank identifier.)
	//for loop with the range keyword
	for _, num := range nums {
		Sum += num
	}

	fmt.Println(Sum)
}
