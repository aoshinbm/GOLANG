package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

// create a method
// in the method the receiver is a pointer to the instance of circle struct
func (c *Circle) calArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	//create a structure variable of circle called c
	c := Circle{radius: 7}

	//use method
	c.calArea()

	fmt.Printf("%+v", c)
}
