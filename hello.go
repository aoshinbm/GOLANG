package main

import "fmt"

func main() {

	//passing a string to print
	fmt.Println("Hello, World!")

	//print a variable
	//declared n initialised a variable
	var city string = "Mumbai"
	//print the variable
	fmt.Print(city)
	fmt.Print("\n")

	//to print variable n string together
	var fname string = "Aoshin"
	var lname string = "Manjuran"
	//printing variable alongwith with some string
	fmt.Print("Welcome ", fname, " ", lname)

	//PROBLEM with PRINT() method is that it wont print new line
	//therefore we can write in this way it will still print together
	var user string = "Aoshu"
	var name string = "Manjn"
	fmt.Print(user)
	fmt.Print(name)
	/*
		OUTPUT :
				AoshuManjn
	*/

	//to print variable using \n(newline character)
	var nam1 string = "Lucian"
	var nam2 string = "Roxane"
	fmt.Print("\n")
	fmt.Print(nam1, "\n")
	fmt.Print(nam2)
	//without even using \n we can print variable on new line using PRINTLN() method
	fmt.Println(nam1)
	fmt.Println(nam2)

	/* rather than making up statement using printf function can be used
	that implements formatted input-output.*/

}
