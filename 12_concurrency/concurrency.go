package main

import (
	"sync"
	"time"
)

// Lets say there are 3 goroutines (f1, f2, f3) f1 and f2 want to send data to f3.
// And all the goroutines want to send the data to main goroutine.

// func f1(ch chan<- int, ch2 chan<- int) { // chan<- int means that this function can only send data to the channel, it cannot receive data from the channel.
// 	ch <- 1
// 	ch2 <- 1 * 2
// }

// func f2(ch chan<- int, ch2 chan<- int) { // chan<- int means that this function can only send data to the channel, it cannot receive data from the channel.
// 	ch <- 2 // Sending 2 to that channal
// 	ch2 <- 2 * 2
// }

// func f3(ch <-chan int, ch2 chan<- int) { // <- chan int means that this function can only receive data from the channel, it cannot send data to the channel.
// 	a := <-ch
// 	b := <-ch

// 	ch2 <- (a + b) * 2
// }

// func main() {
// 	// Channals are used to communicate between goroutines. They can be used to send and receive values of a specific type.
// 	// This are FIFO queues.

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go f1(ch1, ch2) // This will run f1 in a separate goroutine, and it will send 1 to ch1.
// 	go f2(ch1, ch2) // This will run f2 in a separate goroutine, and it will send 2 to ch1.
// 	go f3(ch1, ch2) // This will run f3 in the main goroutine, and it will receive 1 and 2 from ch1.

// 	// <- , this is a blocking operation, it will wait until it receives a value from the channel. So, the main goroutine will wait until it receives values from ch1 and ch2.

// 	a, b, c := <-ch2, <-ch2, <-ch2 // This will receive 1 and 2 from ch1, and it will receive 6 from ch2.

// 	println(a, b, c) // This will print 2, 4, and 6. Order may vary because of the concurrency, but the values will be the same.
// }

// ---------------------------------------------SELECT-----------------------------------------------------------------

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		ch1 <- 1
// 	}()

// 	go func() {
// 		ch2 <- 2
// 	}()

// 	// Using SELECT we can wait for multiple channel operations, and it will execute the case that is ready first.
// 	// If both channels are ready, it will randomly select one of the cases to execute.
// 	select {
// 	case msg1 := <-ch1:
// 		println("Received from ch1:", msg1)
// 	case msg2 := <-ch2:
// 		println("Received from ch2:", msg2)
// 		// default:
// 		// 	println("No channel is ready")
// 	}

// 	// Default - If none of the channels are ready, it will execute the default case.
// 	// This is useful to avoid blocking the main goroutine if none of the channels are ready.

// 	// The output of this program may vary because of the concurrency.
// 	// It may print "Received from ch1: 1" or "Received from ch2: 2" depending on which channel is ready first.
// 	// If we run this with default case, it may print "No channel is ready"
// }

// ---------------------------------------------UnBUFFERED CHANNELS-----------------------------------------------------------------

// ch := make(chan int) // This is an unbuffered channel, it can only hold one value at a time.
// This makes a sync communication.

// Lets say f1 want to send data to f2
// But the channel is not empty, so f1 will block until f2 receives the data from the channel.
// And f2 will block until it receives the data from the channel.

// ----------------------------------------------BUFFERED CHANNELS-----------------------------------------------------------------

// ch := make(chan int, 2) // This is a buffered channel, it can hold 2 values at a time.
// This makes an async communication.

// Lets say f1 want to send data to f2
// But the channel is not empty, so f1 will block until f2 receives the data from the channel.
// And f2 will block until it receives the data from the channel.

// But if we use buffered channel, f1 can send data to the channel without blocking, and f2 can receive data from the channel without blocking, as long as there is space in the buffer.

// ----------------------------------------------WAIT GROUP ------------------------------------------------------

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		time.Sleep(3 * time.Second)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		time.Sleep(5 * time.Second)
	}(&wg)

	wg.Wait() // Execution blocks until the 2 child goorutines completes its execution

}
