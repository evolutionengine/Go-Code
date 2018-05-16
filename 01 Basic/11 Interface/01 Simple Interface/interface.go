package main

import (
	"fmt"
	"math"
)

// Vertex - type struct, contains X & Y
type Vertex struct {
	X, Y float64
}

// Abs - receiver function for vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// altAbs - alternative syntax
func altAbs(v *Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // Note Abs() gets a pointer receiver

	v1 := &Vertex{3, 4}
	fmt.Println(altAbs(v1))
}
