package forloop

import (
	"fmt"
)

func Forloop() {
	c := make(chan int, 1)

	go squares(c)

	for squaredVal := range c {
		fmt.Printf("The squared value is: %d\n", squaredVal)
	}
}

func squares(c chan<- int) {
	for i := 1; i <= 20; i++ {
		c <- i
		fmt.Printf("Len: %d, cap: %d, val: %d\n", len(c), cap(c), i)
	}

	close(c)
}
