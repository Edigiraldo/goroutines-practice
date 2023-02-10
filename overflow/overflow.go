package overflow

import (
	"fmt"
)

func Overflow() {
	fmt.Println("The program started...")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println("The program finished...")
}

func squares(c <-chan int) {
	for i := 0; i <= 4; i++ {
		fmt.Println("The channel value is", <-c)
	}

}
