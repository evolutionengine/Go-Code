package main

import "fmt"

type I interface {
	M() // all receiver functions implementing M() are accessable
}

type T struct {
	S string
}

// M - T implements interface I
// 'I' gets access to 'T' and vice-versa
func (v *T) M() {
	fmt.Println(v.S)
}

func main() {
	var i I = &T{"Hello World"}
	i.M()
}
