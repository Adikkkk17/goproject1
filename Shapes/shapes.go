package Shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}
func (s Square) Perimeter() float64 {
	return 4 * s.Length
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	p := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(p * (p - t.SideA) * (p - t.SideB) * (p - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
func PrintShapeDetails(s Shape) {
	fmt.Printf("Фигура: %T\n", s)
	fmt.Printf("  Площадь: %.2f\n", s.Area())
	fmt.Printf("  Периметр: %.2f\n", s.Perimeter())
	fmt.Println()
}
