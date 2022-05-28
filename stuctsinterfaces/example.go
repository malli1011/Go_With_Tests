package stuctsinterfaces

import (
	"math"
)

type Area float64

type Shape interface {
	getArea() Area
}

type Circle struct {
	Radius float64
}

func (c Circle) getArea() Area {
	temp := math.Pow(c.Radius, 2)
	return Area(math.Pi * temp)
}

type Rectangel struct {
	Breadth float64
	Length  float64
}

func (r Rectangel) getArea() Area {

	return Area(r.Length * r.Breadth)
}

type Square struct {
	Length float64
}

func (s Square) getArea() Area {

	return Area(math.Pow(s.Length, 2.0))
}
