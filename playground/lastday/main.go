package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// n := time.Now()
	newDay := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)
	timeToDoSomething(newDay)
}

func timeToDoSomething(day time.Time) (b bool) {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("[timeToGenerateReport] LoadLocation:", err)
		return
	}
	now := day.In(location)
	lastDay := now.AddDate(0, 0, -1)
	fmt.Println("now", now)
	fmt.Println("lastDay", lastDay)
	targetStart := time.Date(lastDay.Year(), lastDay.Month(), 1, 0, 0, 0, 0, location)
	targetEnd := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 0, 0, 0, 0, location)
	b = lastDay.After(targetStart) && lastDay.Before(targetEnd)
	if os.Getenv("LEVPAY_TEST") == "true" {
		b = true
	}
	fmt.Println("[targetStart]: ", targetStart)
	fmt.Println("[targetEnd]: ", targetEnd)
	return
}
