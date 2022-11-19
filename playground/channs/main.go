package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)

	go listenDone2(ctx)
	go listenDone(ctx)

	time.Sleep(time.Second * 10)
	cancel()
	fmt.Println("cancelled")
	time.Sleep(time.Second * 10)
}

// this fails the test
func listenDone(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("1")
}

// this kills the goroutine
func listenDone2(ctx context.Context) {
	ctxD := ctx.Done()
	for ctxD != nil {
		_, ok := <-ctxD
		if !ok {
			ctxD = nil
			continue
		}
		// select {
		// case :
		// }
	}
	fmt.Println("2")
}
