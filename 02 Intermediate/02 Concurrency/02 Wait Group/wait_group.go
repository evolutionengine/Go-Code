package main

import (
	"sync"
	"fmt"
)

//In order for main() to wait till foo & chu finish execution, we use a package called 'Sync'
//Declare a var of type sync.WaitGroup
//WaitGroup helps in telling main() when other goroutines finish execution
var wg sync.WaitGroup

func foo() {
	//we defer wg.Done(), so that foo signals when its done executing
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Foo is:", i)
	}
}

func chu()  {
	//we defer wg.Done(), so that chu signals when its done executing
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Chu is:", i)
	}
}

func main() {
	//We tell the main() for how many goroutine it has to watch out for
	wg.Add(2)
	//run the goroutines
	go foo()
	go chu()
	//we tell the main() till all the goroutines are done executing
	wg.Wait()
}

//When you run the program you can see the output of foo & chu
//However, the output is random, you can see some 'foo' followed by 'chu' and vice versa
//its because foo & chu are running concurrently, and you don't know the order of execution