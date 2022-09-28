package tutorial

import (
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Add() {
	v.X += 5
	v.Y += 5
}

func (v Vertex) Sub() {
	v.X -= 3
	v.Y -= 3
}
