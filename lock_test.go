package redis

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	ctx := context.Background()
	LockName := "schedule"
	const C = 5
	wg := sync.WaitGroup{}
	wg.Add(C)
	for i := 0; i < C; i++ {
		go func(i int) { // Use multiple goroutines on a single machine to simulate multiple processes in a distributed environment
			defer wg.Done()
			if TryLock(ctx, client, LockName, 10*time.Minute) {
				fmt.Printf("goroutine %d acquired the lock\n", i)
			}
		}(i)
	}
	wg.Wait()
	ReleaseLock(ctx, client, LockName)
}

// go test -v . -run='^TestLock$' -count=1
