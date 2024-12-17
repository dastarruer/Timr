package main

import (
	"fmt"
	"os"
	"time"
	"timr/timer"
	"strconv"
)

func main() {
	args := os.Args
	validNumArgs := 2

	if len(args) != validNumArgs {
		fmt.Println("usage: ./main.go [hours:minutes:seconds]")
		os.Exit(1)
	}

	// Create a ticker; this is a channel that is given a value every time.Second
	durationSeconds, err := strconv.Atoi(args[1])

	// If the user provides an invalid time or mnumber of arguments
	if err != nil {
		fmt.Println("usage: ./main.go [hours:minutes:seconds]")
		os.Exit(1)
	}

	timer := timer.CreateTimer(durationSeconds)

	// Start the timer
	go timer.StartTimer()

	// Stop the main function while the timer runs
	time.Sleep(time.Duration(durationSeconds) * time.Second)

	// Stop the timer
	timer.StopTimer()
	fmt.Println("Time is up!")
}
