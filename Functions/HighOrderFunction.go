package main

import "fmt"

func printResult(radius float64, calculFunct func(r float64) float64) {
	//call the function as input
	result := calculFunct(radius)   //store output of function in the variable
	fmt.Println("Result :", result) //print the result
}
func getFunct(query int) func(r float64) float64 {
	//return a function that takes float as iput and give output
	//map that maps query numbers to function
	query_to_funct := map[int]func(r float64) float64{
		//key datatype=>integer
		//value datatype=>function
		1: calculArea,
		2: calculPerimeter,
		3: calculDiameter,
	}
	return query_to_funct[query]
	//indexing map with query number
}
func main() {
	var query int
	var radius float64
	fmt.Println("Enter the radius of the circle")
	fmt.Scanf("%f", &radius)
	fmt.Println("\nEnter \n 1.Area \n 2.Perimeter \n 3.Diameter ")
	fmt.Scanf("%v", &query)

	printResult(radius, getFunct(query))
}
