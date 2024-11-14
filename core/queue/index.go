package queue

import (
	"sync"
)

func InitQueue() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		QueueMp4Quantity()
	}()

	wg.Wait()
}
