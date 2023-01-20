package main

import "fmt"

func printName(strr string) {
	fmt.Println(strr)
}
func printRollNo(rno int) {
	fmt.Println(rno)
}
func printAddress(addr string) {
	fmt.Println(addr)
}
func main() {

	printName("Marcello")
	defer printRollNo(23)
	//defer will delay execution of printRollNo function
	//untill surrounding function returns
	printAddress("Goregaon")

	/*
		-First control goes to printName function
		-den next it goes to printRollNo but it has defer so v hav to wait
		untill other function is executed
		-therefore control goes to printAddress function which is executed den
		it goes to printRollNo
	*/

}
