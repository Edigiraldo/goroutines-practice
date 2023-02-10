package mutex

import (
	"fmt"
	"sync"
	"time"
)

func Mutex() {
	var counter int
	mutex := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go incrementer(&counter, mutex)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("The value of the counter after being incremented by 1000 goroutines is: %d\n", counter)
	if counter == 1000 {
		fmt.Println(">         -----------------------------------------------------")
		fmt.Println("Race conditions where avoided because of the use of the mutex!")
	}
}

func incrementer(counterPtr *int, m *sync.Mutex) {
	m.Lock()
	*counterPtr += 1
	m.Unlock()
}
