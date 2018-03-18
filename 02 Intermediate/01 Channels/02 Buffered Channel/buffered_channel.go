package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
//Buffered channels are asynchronous

func generate(ch chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 0; i < 100 ; i++ {
		ch <- i
	}

}

func receive(ch chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("Printing numbers")
	var receiver int
		for {
			receiver = <- ch
			fmt.Printf("Received: %v\n", receiver)
		}
}

var wg sync.WaitGroup

func main() {

	//Specifying a value indicates the buffer
	ch := make(chan int, 100)

	//run the goroutines

		wg.Add(2)
		go generate(ch, &wg)
		go receive(ch, &wg)
		wg.Wait()

	fmt.Println("Operation Completed!")
}
