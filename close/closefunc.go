package closefunc

import (
	"fmt"
	"time"
)

func CloseFunc() {
	name := make(chan string, 5)

	go greet(name)

	// name <- "Edison"
	// name <- "Giraldo"
	// name <- "Aristizábal"
	close(name)

	time.Sleep(1 * time.Second)

	fmt.Println("The main program has been stopped...")
}

func greet(name chan string) {
	for i := 0; i < 5; i++ {
		n, ok := <-name
		fmt.Printf("Hello %s, welcome to Apple\tchannel: %t\n", n, ok)
	}
}
