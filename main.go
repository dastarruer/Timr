package main

import (
	"fmt"
	"time"
	"timr/timer"
)



func main() {
	// Create a ticker; this is a channel that is given a value every time.Second
	durationSeconds := 2
	timer := timer.createTimer(durationSeconds)

	// Start the timer
	go timer.startTimer(timer, &durationSeconds)

	// Stop the main function while the timer runs
	time.Sleep(time.Duration(durationSeconds) * time.Second)

	// Stop the timer
	timer.stopTimer(timer)
	fmt.Println("Time is up!")
}




