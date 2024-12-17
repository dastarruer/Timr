package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
	"timr/timer"
)

func main() {
	args := os.Args[1:]

	pattern := `^([1-9]|1[0-9]|2[0-3]):([1-9]|[1-5][0-9]):([1-9]|[1-5][0-9])$`

	re = regexp.MustCompile(pattern)

	// Create a ticker; this is a channel that is given a value every time.Second
	durationSeconds := 70
	timer := timer.CreateTimer(durationSeconds)

	// Start the timer
	go timer.StartTimer()

	// Stop the main function while the timer runs
	time.Sleep(time.Duration(durationSeconds) * time.Second)

	// Stop the timer
	timer.StopTimer()
	fmt.Println("Time is up!")
}
