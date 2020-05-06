package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var iter int
	var wg sync.WaitGroup
	for {
		iter++
		pp := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}
		semaphore := make(chan int, 20)
		for i, p := range pp {
			wg.Add(1)
			semaphore <- 1
			go func(p string, i, it int) {
				defer wg.Done()
				defer func() { <-semaphore }()
				http.Get("https://www.google.com")
				fmt.Printf("iter: %v	i: %v\n", it, i)
			}(p, i, iter)
		}
		wg.Wait()
		close(semaphore)
		<-time.After(100 * time.Millisecond)
	}
}
