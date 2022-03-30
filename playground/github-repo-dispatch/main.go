package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

var eventChan = make(chan string, 100)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go sendEvent(&wg)
	go consumeEvent(&wg)
	wg.Wait()
}

func dispatch() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("TOKEN")})

	tc := oauth2.NewClient(ctx, ts)
	tc.Timeout = time.Second * 10

	client := github.NewClient(tc)

	msg := json.RawMessage(`{"example":"message"}`)

	_, resp, err := client.Repositories.Dispatch(
		context.Background(), "arxdsilva", "golang-ifood-sdk",
		github.DispatchRequestOptions{
			EventType:     "trigger-test",
			ClientPayload: &msg})
	if err != nil {
		fmt.Printf("err: %v, status: %v", err.Error(), resp.Status)
	}
}

func sendEvent(wg *sync.WaitGroup) {
	defer wg.Done()
	var count = 1
	for i := 0; i < count; i++ {
		fmt.Println("sending msg")
		eventChan <- "event"
		time.Sleep(time.Second * 3)
	}
	time.Sleep(time.Second * 5)
}

func consumeEvent(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case msg := <-eventChan:
			fmt.Println("msg received: ", msg)
			go func() {
				time.Sleep(time.Minute)
				dispatch()
				fmt.Println("dispatched")
			}()
		default:
			// fmt.Println("msg not received, waiting 2s")
			time.Sleep(time.Second * 2)
		}
	}
}
