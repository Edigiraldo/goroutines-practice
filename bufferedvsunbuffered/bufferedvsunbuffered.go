package bufferedvsunbuffered

import (
	"fmt"
)

func BufferedVsunbuffered() {
	fmt.Println("The program started...")
	c := make(chan int, 1)
	fmt.Printf("The channel has length: %d and capacity: %d\n", len(c), cap(c))

	go printer(c)

	c <- 1

	fmt.Println("The program finished...")
}

func printer(c chan int) {
	fmt.Println(<-c)
}
