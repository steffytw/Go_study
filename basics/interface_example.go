package main

import "fmt"

type Square struct {
	Side int
}
type Rectangle struct {
	Length  int
	breadth int
}

type Shape interface {
	Area() int
	Perimeter() int
}

func (s Square) Area() int {
	return s.Side * s.Side
}
func (s Square) Perimeter() int {
	return 4 * s.Side
}

func (r Rectangle) Area() int {
	return r.Length * r.GetBreadth()
}
func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.GetBreadth())
}

func (rec *Rectangle) GetBreadth() int {
	return rec.breadth
}

func PrintShape(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

func NewRectangle(length, breadth int) Rectangle {
	return Rectangle{
		Length:  length,
		breadth: breadth,
	}
}

func main() {
	s := Square{
		Side: 10,
	}

	r := NewRectangle(10, 20)

	PrintShape(s)
	PrintShape(r)
}
