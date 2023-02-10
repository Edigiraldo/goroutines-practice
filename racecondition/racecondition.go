package racecondition

import (
	"fmt"
	"time"
)

func RaceCondition() {
	var counter int

	for i := 0; i < 1000; i++ {
		go incrementer(&counter)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("The value of the counter after being incremented by 1000 goroutines is: %d\n", counter)
	if counter != 1000 {
		fmt.Println(">         -----------------------------------------------------")
		fmt.Println("As you can see the values are not the same. This is caused by the race condition of multiple goroutines manipulating the same variable")
	}
}

func incrementer(counterPtr *int) {
	*counterPtr += 1
}
