package main

import "fmt"
import "sync"
import "time"

// MUTEX
/*
	When multiple goroutines (threads) work on the same resource, it can cause uncertain results.
	In the example below, 100 goroutines are utilised to increment a counter value to 100.
	Multiple goroutines accessing and modifying shared data simultaneously can lead to race conditions. 
	Since counter.value++ is not an atomic operation (it performs a read, increment, and write), 
	two goroutines may read the same value before either writes it back, causing updates to be lost. 
	The final counter value therefore becomes unpredictable.

	A mutex is synonymous to a lock. It is used to lock a resource for use by a particular goroutine.
	By locking the shared resource before modifying it and unlocking it afterward, 
	we ensure that only one goroutine updates the counter at a time, preventing race conditions.
*/

type Counter struct {
	value int
	lock sync.Mutex // type of a mutex is sync.Mutex
}

func countWithoutLock(counter *Counter){
	counter.value++
	fmt.Println(counter.value)
}

func countWithLock(counter *Counter){
	counter.lock.Lock() // lock resource
	defer counter.lock.Unlock() // ensure resource is unlocked at the end (even if a panic occurs)
	counter.value++
	fmt.Println(counter.value)
}

// WAIT GROUP
/*
	It can be used to wait for a certain number of tasks/goroutines to complete before moving forward.
	It can be used instead of time.Sleep
*/

func countWithLockAndWaitGroup(counter *Counter, wg *sync.WaitGroup){
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++
	fmt.Println(counter.value)
	wg.Done() // mark the completion of the current goroutine in the wait group
}

func Goroutine3() {
	// Trying to count till 100 using 100 different goroutines

	counter := Counter{value:0}

	// counting without lock

	for i:=0; i<100; i++ {
		go countWithoutLock(&counter) // -> prints 1 to 100 in no specific order
	}

	time.Sleep(2 * time.Second)

	fmt.Println("----")
	counter.value = 0 // resetting the counter value

	// counting with lock

	for i:=0; i<100; i++ {
		go countWithLock(&counter) // -> prints 1 to 100 in correct order
	}

	time.Sleep(2 * time.Second)

	fmt.Println("----")
	counter.value = 0 // resetting the counter value

	// using a wait group to wait for completion of all goroutines instead of time.Sleep

	wg := sync.WaitGroup{}
	wg.Add(100) // wait for 100 tasks/goroutines to finish

	for i:=0; i<100; i++ {
		/* wg.Add(1) 
			-> instead of wg.Add(100), we can add 1 task to wait for in every iteration.
			Just need to make sure that this task to wait for is added in the wait group before the goroutine is started.
		*/
		go countWithLockAndWaitGroup(&counter, &wg)
	}

	wg.Wait() // this will block until all the 100 goroutines are marked done
}