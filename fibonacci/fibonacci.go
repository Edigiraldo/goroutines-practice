package fibonacci

import "fmt"

func Fibonacci(num int) {
	c := make(chan int)

	fmt.Printf("These are the first %d numbers in the fibonacci sequence\n", num)

	go func(c chan<- int) {
		counter := 0
		for i, j := 0, 1; counter < num; i, j = i+j, i {
			c <- i
			counter++
		}

		close(c)
	}(c)

	counter := 0
	for val := range c {
		fmt.Printf("Fibo[%d]: %d\n", counter, val)
		counter++
	}
}
