package main

import (
	// "fmt"
	// "time"
	// "timr/timer"
	"timr/tui"
)

func main() {
	// Create a ticker; this is a channel that is given a value every time.Second
	// durationSeconds := 1
	// timer := timer.CreateTimer(durationSeconds)

	// // Start the timer
	// go timer.StartTimer()

	// // Stop the main function while the timer runs
	// time.Sleep(time.Duration(durationSeconds) * time.Second)

	// // Stop the timer
	// timer.StopTimer()
	// fmt.Println("Time is up!")

	tui.HelloWorld()
}
