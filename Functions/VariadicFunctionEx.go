package main

import "fmt"

func printDetails(Student string, subjects ...string) {
	fmt.Println("Hey", Student, "here are ur subjects: ")
	for _, subj := range subjects {
		fmt.Printf("%s", subj)
	}
}

func main() {
	printDetails("Joey")
	printDetails("Rachel", "Biology", "Chemistry")
	fmt.Println("")
	printDetails("Lucian", "Maths", "Biology", "Chemistry")
}
