package main

import (
	"fmt"
	"time"
)

// DIRECTIONAL CHANNELS

// Send only channel
func send(ch chan<- int){
	fmt.Println("Sent to channel: 20")
	// a := <- ch // Error: invalid operation: cannot receive from send-only channel chan<- int ch (variable of type chan<- int)
	ch <- 20
}

// Receive only channel
func receive(ch <-chan int) {
	a := <- ch
	// ch <- 20 // Error: invalid operation: cannot send to receive-only channel <-chan int ch (variable of type <-chan int)
	fmt.Println("Received from channel:",a)
}

// BUFFERED CHANNELS

/*
	In a normal channel, the program blocks after sending one value through the channel for it to be received.
	With a buffered channel, the program does not block. The value goes to the buffer and can be received later.
*/

func sendThreeValues(ch chan int){
	ch <- 10
	ch <- 20
	ch <- 30
}

func Goroutine2() {
	// // Using unidirectional channels to send and receive values

	// ch := make(chan int) // bidirectional channel

	// /*
	// 	we can explicitly type cast the bidirectional channel to only send or receive and then pass to the corresponding functions as below:
	// 	sendCh := (chan<- int)(ch)      
	// 	receiveCh := (<-chan int)(ch) 

	// 	However, this is not required in practice as the Go compiler can perform implicit conversion at compile time.
	// 	So, passing a bidirectional channel to a function accepting unidirectional channel works and is the intended usage.
	// */

	// go send(ch)
	// go receive(ch)

	time.Sleep(2 * time.Second) // needed to stop the main thread from terminating before the goroutines finish

	// // using buffered channels to send values without blocking

	// /*
	// 	The below non-buffer channel example causes the deadlock error, because it is sending a value in the channel, then
	// 	waits fot it to be received and never runs the next line. 

	// 	ch1 := make(chan int)
	// 	ch1 <- 10
	// 	fmt.Println("this never runs")
	// 	val0 := <- ch1

	// 	However, using a buffered channel here can solve this problem as the program does not wait for the value to be received,
	// 	instead the value is passed to a buffer and received later.
	// */
	// ch1 := make(chan int, 1) // buffered channel that can hold 1 value
	// ch1 <- 10
	// fmt.Println("this runs")
	// val0 := <- ch1
	// fmt.Println("value received from buffer:", val0)


	// // Sending less or equal values than the buffer size - never blocks

	// ch2a := make(chan int, 3) // buffered channel that can hold 3 values
	// ch2a <- 10
	// ch2a <- 20
	// ch2a <- 30

	// fmt.Println("even though values from ch2a are never received, this will not block")

	// ch2 := make(chan int, 3)
	// go sendThreeValues(ch2)

	// val1 := <- ch2
	// val2 := <- ch2
	// val3 := <- ch2
	// // Since the channel can only send 3 values, the below line causes a deadlock error as it is waiting to receive a value that is never sent
	// // val4 := <- ch2 // -> Error: fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(val1, val2, val3)

	// // Sending more values than the buffer size - can cause blocking

	// /* 
	// 	The below two examples does the same thing: send 3 values in a channel and receive only 2.
	// 	However, the first example does not cause a deadlock issue, but the second one does. 
	// 	This is because sending more values in a channel than being received is a blocking operation.
	// 	However, the difference is that the first example is blocking a separate thread created via goroutine,
	// 	whereas, the second example is blocking the main thread which is throwing the deadlock error.
	// */

	// // first example - blocking goroutine thread
	// ch3 := make(chan int, 2)
	// go sendThreeValues(ch3)

	// val4 := <- ch3
	// val5 := <- ch3
	// fmt.Println(val4, val5) // this works and values are printed.

	// // second example - blocking main thread
	// // ch4 := make(chan int, 2)
	// // ch4 <- 10
	// // ch4 <- 20
	// // ch4 <- 30
	
	// // val6 := <- ch4
	// // val7 := <- ch4

	// // fmt.Println(val6, val7) // this never runs as thread gets blocked waiting for value to be received

	// /*
	// 	This can be handled by making space in the buffer for the extra value by receiving some values
	// */

	// ch4 := make(chan int, 2)
	// ch4 <- 10
	// ch4 <- 20

	// val6 := <- ch4 // clearing out space for a third value by receiving the first value
	// fmt.Println("val6:", val6)

	// ch4 <- 30
	
	// val7 := <- ch4
	// val8 := <- ch4
	// fmt.Println("val7 and val8:", val7, val8)
	
	Goroutine3()
} 