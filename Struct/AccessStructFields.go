package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius int
}

func main() {
	var c Circle //struct variable to acess struct fields
	c.radius = 8
	c.x = 5
	c.y = 7
	fmt.Printf("%+v \n", c)
}
