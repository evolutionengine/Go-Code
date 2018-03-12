package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var (
	counter int
	wg sync.WaitGroup
	//Here we use sync.Mutex , with this the variable is locked till the operation is executed
	//After execution, it is unlocked for other operations to take over
	mutex sync.Mutex
)

func increment(s string)  {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Printf("%s: %v\n", s, counter)
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
