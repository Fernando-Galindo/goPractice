package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length int
	width  int
}
type CIRCLE struct {
	radius float64
}

func (s SQUARE) Area() int {
	return s.length * s.width
}
func (c CIRCLE) Area() int {
	return int(math.Pow(c.radius, 2) * math.Pi)
}

type SHAPE interface {
	Area() int
}

func info(s SHAPE) {
	fmt.Println(s.Area())
}

func main() {
	s := SQUARE{
		length: 5,
		width:  8,
	}
	c := CIRCLE{
		radius: 4,
	}

	info(s)
	info(c)
}
