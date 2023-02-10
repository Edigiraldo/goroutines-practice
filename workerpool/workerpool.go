package workerpool

import (
	"fmt"
	"time"
)

func Workerpool() {
	valuesToPrint := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	nWorkers := 3
	jobsChannel := make(chan int)

	for workerID := 0; workerID < nWorkers; workerID++ {
		go worker(jobsChannel, workerID)
	}

	for _, val := range valuesToPrint {
		jobsChannel <- val
	}

	close(jobsChannel)

	// Waiting for workers to finish last jobs
	// time.Sleep(4 * time.Millisecond)
}

func worker(jobsChannel <-chan int, workerID int) {
	for val := range jobsChannel {
		fmt.Printf("I'm the worker %d an am printing the value %d\n", workerID, val)
		time.Sleep(10 * time.Millisecond)
	}
}
