package main

import "fmt"

func operations(a int, b int) (sum int, subt int) {
	//sum & subt are named return parameters
	//return values names have been declared & initialized in function signature
	sum = a + b
	subt = a - b
	return
}

func main() {

	sum, diff := operations(100, 66)
	fmt.Println(sum, diff)
}
