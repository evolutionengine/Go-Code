package main

import (
	"sync"
	"fmt"
	"runtime"
)

//This is exact same e.g as previous except we are adding parallelism
//We explicitly tell the program to use all the cores or cpu during runtime,
//this is exactly where simplicity & ease of concurrency features shine in Golang
//It just makes it so easy to have concurrent threads along with parallel execution
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var wg sync.WaitGroup

func foo() {

	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Foo is:", i)
	}
}

func chu()  {

	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Chu is:", i)
	}
}

func main() {

	wg.Add(2)
	go foo()
	go chu()
	wg.Wait()
}
