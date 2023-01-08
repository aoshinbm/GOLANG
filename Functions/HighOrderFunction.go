package main

import "fmt"

func printResult(radius float64, calculFunct func(r float64) float64) {
	result := calculFunct(radius)
	fmt.Println("Result :", result)
}
func getFunct(query int) func (r float64)float64  {
	query_to_funct :=map[int]func (r float64)float64  {
		1:calculArea
		2:calculPerimeter
		3:calculDiameter
	}
	return query_to_funct[query]
}  {
	
}