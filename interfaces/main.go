package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

type Square struct {
	side float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r // pi*radius*radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r // 2*pi*radius
}

func (r Square) Area() float64 {
	return r.side * r.side // 2*side
}

func (r Square) Perimeter() float64 {
	return 2 * r.side // 2*side
}

func main() {
	var c Circle
	c.r = 2.5
	fmt.Println("Area of circle is ", c.Area())
	fmt.Println("Perimeter of circle is ", c.Perimeter())
	var sq Square
	sq.side = 4.0
	fmt.Println("Area of square is ", sq.Area())
	fmt.Println("Perimeter of Square is ", sq.Perimeter())

}
