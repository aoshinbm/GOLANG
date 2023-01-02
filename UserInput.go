package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter name : ")
	//format specifier %s is for string
	//entered value is stored into the name variable
	fmt.Scanf("%s", &name)
	fmt.Println("Hey ", name, "!! Wassup ?")

	//taking multiple user inputs
	var user string
	var is_muggle bool
	fmt.Println("Enter name & are you muggle : ")
	fmt.Scanf("%s %t", &user, &is_muggle)

	fmt.Println(user, is_muggle)

	var a string
	var b int
	fmt.Println("Enter a string & a number :")
	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count :", count)
	fmt.Println("err :", err)
	fmt.Println("a :", a)
	fmt.Println("b :", b)

}
