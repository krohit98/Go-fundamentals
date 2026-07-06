// README: run one example at a time to avoid confusion

package main

import "fmt"
import "time"

func run1(){
	h, m, s := time.Now().Clock()
	time.Sleep(2 * time.Second) // delay 2s
	fmt.Printf("run1 started at: %02d:%02d:%02d\n", h, m, s)
}

func run2(){
	h, m, s := time.Now().Clock()
	time.Sleep(2 * time.Second) // delay 2s
	fmt.Printf("run2 started at: %02d:%02d:%02d\n", h, m, s)
}

func run3(){
	h, m, s := time.Now().Clock()
	time.Sleep(2 * time.Second) // delay 2s
	fmt.Printf("run3 started at: %02d:%02d:%02d\n", h, m, s)
}

// Returning from a function run on a separate thread using channel
func addNums(x int, y int, ch chan int){
	time.Sleep(2 * time.Second)
	fmt.Println("adding",x,"and",y)
	ch <- x + y // to return a value through a channel use <-
} 

// this function is expected to return a value through the channel, but never does - this causes a deadlock error
func addNumsWithoutReturn(x int, y int, ch chan int){
	time.Sleep(2 * time.Second)
	fmt.Println("this function does not return a value in the channel")
}

func Goroutine1() {
	/*
		Goroutines enables developers to create multi threaded programs and run these threads concurrently
	*/

	// Without goroutines - each start after the last one finishes
	run1() // -> run1 started at: 20:12:31
	run2() // -> run2 started at: 20:12:33
	run3() // -> run3 started at: 20:12:35

	// With goroutines - all start at the same time without waiting for the last one
	go run1() // use the go keyword to run a function in a separate thread
	go run2()
	go run3()
	/*
		->
		run2 started at: 20:15:13
		run1 started at: 20:15:13
		run3 started at: 20:15:13

		the order at which the functions are run might not be the same as the order they are called in
	*/

	time.Sleep(4 * time.Second) // need to wait for the goroutines to finish before the program terminates to print in terminal

	// Channels

	/* 
		They unable devs to return values from goroutines
		They create blocking operations
		The program waits for a thread to complete and return the result in a channel before moving forward
	*/

	ch := make(chan int) // create a channel
	go addNums(5, 10, ch)
	res1 := <- ch // this is a blocking operation
	
	fmt.Println(res1)

	// Potential deadlock errors when using channels

	ch2 := make(chan int)
	go addNumsWithoutReturn(2, 10, ch2) // this function does not return any value in the channel
	/*
		The current program thread (main thread) is waiting for the new thread to return a value which it never does.
		This causes a deadlock error because the main thread cannot proceed until the new thread returns a value
		which it never does.
	*/
	res2 := <- ch2 // -> Error (after 2s): fatal error: all goroutines are asleep - deadlock!
	fmt.Println(res2)

	// Multiple goroutines using the same channel - waiting for first value to be returned

	ch3 := make(chan int)
	go addNums(5, 1, ch3)
	go addNums(5, 2, ch3)
	go addNums(5, 3, ch3)
	go addNums(5, 4, ch3)
	/*
		-> adding 5 and 3
	*/

	res2b := <- ch3 // just waits for the first thread to return a value in the channel | the other goroutines will still complete but the program will not wait on this step to catch their returned values
	fmt.Println("res2 value:", res2b) // -> res2b value: 8

	// Multiple goroutines using the same channel - waiting for all values to be returned (in any order)

	ch4 := make(chan int)
	go addNums(5, 1, ch4)
	go addNums(5, 2, ch4)
	go addNums(5, 3, ch4)
	go addNums(5, 4, ch4)
	/*
		->
		adding 5 and 2
		adding 5 and 1
		adding 5 and 3
		adding 5 and 4
	*/

	// waits for all the 4 threads to run and finish in any order, then prints the value returned by the last one.
	res3 := <- ch4
	res3 = <- ch4
	res3 = <- ch4
	res3 = <- ch4

	fmt.Println("res3 value:", res3) // -> res3 value: 9

	// Multiple goroutines using the same channel - waiting for all values to be returned (in correct order)

	ch5 := make(chan int)

	// blocking code that waits for each goroutine to finish before starting the next one - synchronous execution
	go addNums(5, 1, ch5)
	res4 := <- ch5
	go addNums(5, 2, ch5)
	res4 = <- ch5
	go addNums(5, 3, ch5)
	res4 = <- ch5
	go addNums(5, 4, ch5)
	res4 = <- ch5
	/*
		->
		adding 5 and 1
		adding 5 and 2
		adding 5 and 3
		adding 5 and 4

		each printed after 2s
	*/

	// prints the value of the last goroutine
	fmt.Println("res4 value:", res4) // -> res4 value: 9

	// Multiple go routines using different channels - waiting for all the channels to receive values

	ch6 := make(chan int)
	ch7 := make(chan int)

	go addNums(5, 1, ch6)
	go addNums(5, 2, ch7)
	/*
		->
		adding 5 and 2
		adding 5 and 1
	*/

	// even though ch7 got the value before ch6, it could not be printed unless both the values are present
	// using this syntax forces blocking code to get all the values before processing them
	resX := <- ch6
	resY := <- ch7

	fmt.Println(resX, resY) // -> 6 7

	// Multiple go routines using different channels - processing the first value as it comes using select

	ch8 := make(chan int)
	ch9 := make(chan int)

	go addNums(5, 1, ch8)
	go addNums(5, 2, ch9)

	/* -> adding 5 and 2 */

	select {
	case res5 := <- ch8:
		fmt.Println("res5 value:", res5)
	case res6 := <- ch9:
		fmt.Println("res6 value:", res6)
	}
	/* -> res6 value: 7 */

	// Multiple go routines using different channels - processing all the values one by one as they come

	ch10 := make(chan int)
	ch11 := make(chan int)

	go addNums(5, 1, ch10)
	go addNums(5, 2, ch11)

	// The loop will make the select statement run twice and wait to receive a value in the channels
	for i := 0; i < 2; i++ {
		select {
		case res7 := <- ch10:
			fmt.Println("res7 value:", res7)
		case res8 := <- ch11:
			fmt.Println("res8 value:", res8)
		}
	}
	/* 
		-> 
		adding 5 and 2
		res8 value: 7
		adding 5 and 1
		res7 value: 6
	*/

	fmt.Println("Done")

	Goroutine2()
}