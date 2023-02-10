package sendchannel

import (
	"fmt"
	"runtime"
	"sync"
)

var executedGoroutines string = ""

func SendChannel() {
	mutex := &sync.Mutex{}
	numbersToPrint := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	chanOfChans := make(chan chan int)

	for i := 0; i < 2; i++ {
		go anotherGoroutine(chanOfChans, mutex)
		go printFromChan(chanOfChans, mutex)
		go anotherGoroutine(chanOfChans, mutex)
		go printFromChan(chanOfChans, mutex)
		go anotherGoroutine(chanOfChans, mutex)
		go anotherGoroutine(chanOfChans, mutex)
		go anotherGoroutine(chanOfChans, mutex)
		go printFromChan(chanOfChans, mutex)
		go printFromChan(chanOfChans, mutex)
		go anotherGoroutine(chanOfChans, mutex)
	}

	for val := range numbersToPrint {
		channel := make(chan int, 1)

		channel <- val
		chanOfChans <- channel
	}

	close(chanOfChans)
	fmt.Println("Order of execution in Goroutines:", executedGoroutines)
	fmt.Println("Order expected in execution     :", "paapaaappapaapaaappa")

	fmt.Println("The number of threads being executed concurrently was:", runtime.GOMAXPROCS(0))
}

func anotherGoroutine(chanOfChans chan chan int, mutex *sync.Mutex) {
	// fmt.Println("I'm another. Greets!!")
	mutex.Lock()
	executedGoroutines += "a"
	mutex.Unlock()

	<-chanOfChans
}

func printFromChan(chanOfChans chan chan int, mutex *sync.Mutex) {
	// fmt.Println("printer was executed!!")
	// fmt.Printf("chanOfChans(%v) - len(%d) - cap(%d)\n", chanOfChans, len(chanOfChans), cap(chanOfChans))

	mutex.Lock()
	executedGoroutines += "p"
	mutex.Unlock()

	<-chanOfChans
	// chanCounter := 0
	// for c := range chanOfChans {
	// 	fmt.Printf("getting chan!!\n")
	// 	valsCounter := 0
	// 	for value := range c {
	// 		fmt.Printf("Chan %d[%d]: %d\n", chanCounter, valsCounter, value)
	// 	}
	// }
}
