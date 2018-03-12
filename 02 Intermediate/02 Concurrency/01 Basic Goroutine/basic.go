package main

import "fmt"

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo is:", i)
	}
}

func chu() {
	for i := 0; i < 45; i++ {
		fmt.Println("Chu is:", i)
	}
}

func main() {
	//Go routines can be simply declared using the keyword "go" followed by func name
	go foo()
	go chu()
}

//Well....to your surprise, this program doesn't output anything !!

/*
Because -
1) We have three go routines running simultaneously -> main, foo & chu
2) main() finishes execution before foo & chu can execute and exits the program
3) By default main() runs as a go routine, i.e it runs on a separate thread
4) If foo & chu have to print, we have to find a way to tell 'main' to wait till foo & chu finished execution

To know how we can tell main() to wait, check out the next folder 'Wait Group'
*/