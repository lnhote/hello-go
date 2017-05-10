package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	flatInst := MyFloat(-math.Sqrt2)
	vertexInst := Vertex{3, 4}

    a = flatInst
    fmt.Println(a.Abs()) // a MyFloat implements Abser

	a = &vertexInst // a *Vertex implements Abser
    fmt.Println(a.Abs())

	// a = vertexInst // This is WRONG!!! vertexInst is a Vertex (not *Vertex)
    // fmt.Println(a.Abs())
}
