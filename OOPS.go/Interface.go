package main

import "fmt"

type Shape2D interface {
	//created same function for triangle & rectangle
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Triangle struct {
	height float64
	base   float64
	sideA  float64
	sideB  float64
	sideC  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.height * t.base
}
func (t Triangle) Perimeter() float64 {
	return t.sideA + t.sideB + t.sideC
}

func main() {
	rect := Rectangle{2.5, 3.6}
	fmt.Println("Area of Rectangle:", rect.Area())
	fmt.Println("Perimeter of Rectangle:", rect.Perimeter())

	tri := Triangle{2.0, 5.0, 6.0, 3.0, 4.0}
	fmt.Println("Area of Triangle:", tri.Area())
	fmt.Println("Perimeter of Triangle:", tri.Perimeter())

}
