package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (cir Circle) Area() float64 {
	res := math.Pi * math.Pow(cir.Radius, 2)
	return res
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}

func printAreas(shapes []Shape) {
	for _, s := range shapes {
		fmt.Printf("Площадь - %.2f\n", s.Area())

	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 4.0},
		Circle{Radius: 5.0},
		Rectangle{Width: 2.0, Height: 4.0},
		Rectangle{Width: 4.0, Height: 4.0},
	}
	printAreas(shapes)
}
