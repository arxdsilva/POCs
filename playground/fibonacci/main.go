package main

import (
	"fmt"
	"math/big"
	"sync"
)

func Fibonacci(n uint, wg *sync.WaitGroup) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	var n2, n1 = big.NewInt(0), big.NewInt(1)
	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}
	fmt.Print(".")
	wg.Done()
	// fmt.Println(".", n1)
	return n1
}

func main() {
	var wg sync.WaitGroup
	for fib := 1000; fib < 1101; fib++ {
		wg.Add(1)
		go Fibonacci(uint(fib), &wg)
	}
	wg.Wait()
}
