//This is an interesting example of concurrency using channels
//Note that we are not using sync or mutex
package main

import "fmt"

//This generates numbers and put them on the channel
func generate(ch chan int, nums int)  {
	fmt.Println("Starting to generate numbers - ")
	//start generating nums
	for i := 0; i <= nums; i++ {
		//put each number on channel
		ch <- i
	}
	//after all the nums have been generated, we will close the channel
	//otherwise the receiver does not know when to stop, this results
	//in deadlock as generate is not creating values, but receiver is expecting values
	close(ch)
}

//This flushes the channel of data
func receiver(ch chan int, isdone chan bool)  {
	//we loop using range, range also acts like wait till we loop through all the data
	for num := range ch {
		//print nums as we receive them, since the data is transmitted from channel, it acts like FIFO
		//and we print the nums in perfect order
		fmt.Printf("Received No: %v\n", num)
	}
	//after flushing the channel, we signal that the operation is done
	isdone <- true
}

func main() {
	//create a channel of type int to carry our numbers
	//Just for the sake of education, we add a buffer of size 10, it means at a time 10 numbers can be
	//loaded onto the channel
	ch := make(chan int, 10)
	//create a channel of type bool to indicate the status of operation, we use it for stopping the main() from finishing
	done := make(chan bool)

	var nums int
	fmt.Printf("Enter the nos you want to print: ")
	fmt.Scan(&nums)
	//launch the go routines, they run independently of each other and main()
	go generate(ch, nums)
	go receiver(ch, done)
	//we indicate to main() to stop till we get the operation status
	//<- done actually discards the value, here it acts as a mechanism of signal for the main goroutine to stop
	<- done
	fmt.Println("Operation Completed!")
}
