package main

import (
	"io"
	"log"
	"sync"
)

// closing already closed pipes does not do anything
func main() {
	pr, pw := io.Pipe()
	if err := pr.Close(); err != nil {
		log.Println("err: ", err)
	}
	if err := pr.Close(); err != nil {
		log.Println("err: ", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := pw.Close(); err != nil {
			log.Println("err: ", err)
		}
		if err := pw.Close(); err != nil {
			log.Println("err: ", err)
		}
	}()
	wg.Wait()
}
