package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) add() {
	v.X += 5
	v.Y += 5
}

func (v Vertex) sub() {
	v.X -= 3
	v.Y -= 3
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// add function
	v.add()
	fmt.Println(v)
	fmt.Println(v.Abs())
	//subtraction function
	v.sub()
	fmt.Println(v)
	fmt.Println(v.Abs())

}
