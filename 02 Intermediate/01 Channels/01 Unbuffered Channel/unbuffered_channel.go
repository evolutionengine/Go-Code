//Channels are like pipes from which data can be transferred
//They are particularly important for concurrency as it avoids race conditions
package main

import (
	"fmt"

	"time"
)


//By default communication is synchronous & unbuffered
//The go routines would be blocked if there the receiver is not ready
//as unbuffered channels have no provision of containing data
func sendData(ch chan string)  {

	ch <- "Mumbai"
	ch <- "Chennai"
	ch <- "Delhi"
	ch <- "Banglore"
	ch <- "Amsterdam"
}

func receiveData(ch chan string)  {

	var receiver string
	for {
		receiver = <- ch
		if receiver == "Amsterdam" {
			fmt.Println("Yo, you have reached Amstedam")
		} else {
			fmt.Printf("Received: %s\n", receiver)
		}
	}
}

func main() {
	//Here we are declaring and initializing a channel of type string
	ch := make(chan string)

	//We start two go routines, one sends data, other receives the data
	go sendData(ch)
	go receiveData(ch)
	time.Sleep(1e9)
}

//Unbuffered channels makes synchronization easy, some sometimes it may be restrictive
//In such cases we can use buffered channels
