package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// value receiver
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// value receiver
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// take interface of shape
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	//data types that represent a set of method signatures for  structs
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	//1st way
	// fmt.Println(circle.area())

	//2nd way
	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))

}
