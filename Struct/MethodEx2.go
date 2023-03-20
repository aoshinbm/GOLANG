package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

// create a method
// in this case receiver is the instance of circle struct
func (c Circle) calArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	//create a structure variable of circle called c
	c := Circle{radius: 5}

	//use method
	c.calArea()

	fmt.Printf("%+v", c)
}
