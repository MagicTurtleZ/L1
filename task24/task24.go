package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

// структура Point с инкапсулированными параметрами x,y
type Point struct {
	x float64
	y float64
}

// конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// √((x2​−x1​)^2 + (y2​−y1​)^2)
func Interval(a, b *Point) float64 {
	return math.Sqrt(math.Pow(b.x - a.x, 2) + math.Pow(b.y - a.y, 2))
}

func main() {
	a := NewPoint(6, -5)
	b := NewPoint(10, -8)
	fmt.Println(Interval(a, b))
}
