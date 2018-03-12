package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var (
	counter int
	wg sync.WaitGroup
)

func increment(s string)  {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
}

func main() {
	wg.Add(2)
	go increment("Foo:")
	go increment("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

//go run -race race.go -> checks for any race conditions