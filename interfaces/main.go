package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return math.Pi * 2 * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area: %f\n", s.area())
	fmt.Printf("circumf: %f\n", s.circumf())
}

func main() {
	shapes := []shape{
		square{5},
		square{10},
		circle{5},
		circle{10},
	}
	for _, shape := range shapes {
		printShapeInfo(shape)
	}
	var shape shape = square{5}
	fmt.Println(shape.area())
}
