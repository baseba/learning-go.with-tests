package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func main() {
	r1 := Rectangle{10.0, 10.0}
	r2 := Rectangle{12.0, 6.0}
	c1 := Circle{10.0}
	fmt.Println(Perimeter(r1))
	fmt.Println(r2.Area())
	fmt.Println(c1.Area())

}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
