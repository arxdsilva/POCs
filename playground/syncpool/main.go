package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name string
	Age  int
}

var userPool = sync.Pool{
	New: func() any {
		return &User{}
	},
}

type Set map[string]struct{}

func AllocateViaPoolSyncNoSleep(n int) Set {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	res := make(Set)
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()

			// Pool operations don't need mutex - sync.Pool is thread-safe
			q := userPool.Get().(*User)
			defer userPool.Put(q)
			q.Name = "kieran"
			q.Age = 27

			// Only protect the shared map access
			mu.Lock()
			res[fmt.Sprintf("%p", q)] = struct{}{}
			mu.Unlock()
		}()
	}
	wg.Wait()
	return res
}

func AllocateViaPoolSyncWithSleep(n int) Set {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	res := make(Set)
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()

			// Pool operations don't need mutex - sync.Pool is thread-safe
			q := userPool.Get().(*User)
			defer userPool.Put(q)
			time.Sleep(1 * time.Nanosecond) // simulate work
			q.Name = "kieran"
			q.Age = 27

			// Only protect the shared map access
			mu.Lock()
			res[fmt.Sprintf("%p", q)] = struct{}{}
			mu.Unlock()
		}()
	}
	wg.Wait()
	return res
}

func main() {
	pointers := AllocateViaPoolSyncNoSleep(50_000)
	fmt.Printf("Number of pointers in sync pool allocation: %d\n", len(pointers))
}
