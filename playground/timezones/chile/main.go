package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	locChile, err := time.LoadLocation("America/Santiago")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Locale Chile: " + locChile.String())

	locPeru, err := time.LoadLocation("America/Lima")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Locale Peru: " + locPeru.String())
	start := time.Now()
	yesterday := start.Add(-time.Hour * 24)

	startDefaultTimePeru := time.Date(
		yesterday.Year(),
		yesterday.Month(),
		yesterday.Day(),
		00,
		00,
		00,
		00,
		locPeru,
	).Local().Format("2006-01-02 15:04:05")
	fmt.Printf("start chile: %s\n", startDefaultTimePeru)
	endDefaultTimePeru := time.Date(
		yesterday.Year(),
		yesterday.Month(),
		yesterday.Day(),
		23,
		59,
		59,
		99,
		locPeru,
	).Local().Format("2006-01-02 15:04:05")
	fmt.Printf("start chile: %s\n", endDefaultTimePeru)
}
