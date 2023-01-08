package main

import "fmt"

func calculArea(r float64) float64 {
	return 3.14 * r * r
}
func calculPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func calculDiameter(r float64) float64 {
	return 2 * r
}
func main() {
	var query int
	var radius float64
	fmt.Println("Enter radius:")
	fmt.Scanf("%.2f", &radius)
	fmt.Println("Enter query:\n1)Area\n2)Perimeter\n3)Diameter")
	fmt.Scanf("%v", &query)

	if query == 1 {
		fmt.Println(calculArea(radius))
	} else if query == 2 {
		fmt.Println(calculPerimeter(radius))
	} else if query == 3 {
		fmt.Println(calculDiameter(radius))
	} else {
		fmt.Println("Invalid choice")
	}
}
