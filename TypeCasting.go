package main

import "fmt"

func main() {
	//Converting integer into float
	fmt.Println("Converting int into float:")
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f", f)
	fmt.Println("---------------------------------------")

	//Converting float into integer here numbers decimal pt will be lost
	fmt.Println("Converting float into int:")
	var fl float64 = 78.69
	var a int = int(fl)
	fmt.Printf("%v", a)

	fmt.Println("---------------------------------------")

}
