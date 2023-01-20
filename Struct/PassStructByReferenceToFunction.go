package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

//input parameter as pointer to CIRCLE
func calcArea(c *Circle) {
	const PI float64 = 3.142
	var area float64
	area = (PI * c.radius * c.radius)
	//de-refernce pointer n store calculated area in area field of CIRCLE
	(*c).area = area
}
func main() {
	c := Circle{x: 5, y: 5, radius: 5, area: 0}
	fmt.Printf("%+v \n", c)
	calcArea(&c) //passing address of circle variable using address operator
	fmt.Printf("%+v \n", c)

}

/*
This is bcuz variable c is passed as reference to the calcArea function
hence any change made to variable inside function affects the original argument as well.
OUTPUT:-
{x:5 y:5 radius:5 area:0}
{x:5 y:5 radius:5 area:78.55}
*/
