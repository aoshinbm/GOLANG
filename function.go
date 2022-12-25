package main

import "fmt"

func swap(str1, str2 string) (string, string) {
	return str1, str2
}

//function definition
func mulFloatNumbers(numb1, numb2 float32) float32 {
	product := numb1 * numb2
	return product
	//	return numb1*numb2
}

//function definition
func mulFloatNumbers(numb1, numb2 float32) float32 {
	product := numb1 * numb2
	return product
	//	return numb1*numb2
}

//function definition
func multiplyFloat(n1, n2 float32) float32 {
	return n1 * n2
}

func main() {
	lname, fname := swap("Liam", "Farewell")
	fmt.Println(lname, fname)

	//function call
	result := mulFloatNumbers(25.3, 65.2)
	fmt.Println("25.3 X 65.2 = ", result)

	fmt.Println("52.3 X 1.2 = ", multiplyFloat(52.3, 1.2))
}
