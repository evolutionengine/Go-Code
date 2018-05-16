package main

import (
	"fmt"
	"math"
)

// I - Type interfaces, contains M()
type I interface {
	M() // interface 'I' has access to all the methods of M()
	// 'I' can access 'T' as well as 'F', as both of them implement 'I' via M()
}

// T - type struct, contains S -> string
type T struct {
	S string
}

// M - receiver function, implementing interface 'I'
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("%v %T\n", i, i)
}
