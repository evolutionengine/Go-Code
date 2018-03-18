package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
	//Atomic gives even more fine control than mutex, it
	//literally means doing only one operation at a time on variable
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func increment(s string)  {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		//Since we carry out an atomic function, we don't require lock & unlock
		atomic.AddInt64(&counter, 1)
		fmt.Printf("%s: %v\n", s, counter)
	}
}

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
