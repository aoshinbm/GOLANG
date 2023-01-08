package main

import "fmt"

func addNos(a int, b int) string {
	sum := a + b
	return sum
}
func main() {

	sumofNos := addNos(52, 60)
	fmt.Println(sumofNos)
}

/*
bcuz return typt is not matching

OUTPUT:=
cannot use sum (variable of type int) as type string in return statement
*/
