package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Waitgroup() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go service(wg, i)
	}

	wg.Wait()
}

func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Service called on instance %d\n", instance)
	wg.Done()
}
