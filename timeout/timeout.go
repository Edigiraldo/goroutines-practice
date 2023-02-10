package timeout

import (
	"fmt"
	"time"
)

func Timeout() {
	start := time.Now()
	chan1 := make(chan int)
	chan2 := make(chan int)

	go c1(chan1)
	go c2(chan2)

	select {
	case v1 := <-chan1:
		fmt.Printf("The channel 1 sent the value %v\n", v1)
	case v2 := <-chan2:
		fmt.Printf("The channel 2 sent the value %v\n", v2)
	case <-time.After(1 * time.Microsecond):
		fmt.Printf("The channels last more than 1 microseconds and timed out")
	}

	fmt.Printf("Time elapsed %v\n", time.Since(start))
}

func c1(c chan<- int) {
	time.Sleep(time.Millisecond)
	c <- 1000
}

func c2(c chan<- int) {
	time.Sleep(time.Millisecond)
	c <- 2000
}
