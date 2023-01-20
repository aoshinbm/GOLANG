package main

import "fmt"

func swap_nums(n1_ptr, n2_ptr *int) {
	temp := *n1_ptr
	*n1_ptr = *n2_ptr
	*n2_ptr = temp
}
func main() {
	numb1 := 10
	numb2 := 56
	fmt.Println(numb1, numb2)

	swap_nums(&numb1, &numb2)
	fmt.Println(numb1, numb2)
}
