package timer

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Timer struct {
	ticker          *time.Ticker
	done            chan bool
	durationSeconds int
}

// Only works for linux/mac
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Return a pointer to a new Timer instance
func CreateTimer(durationSeconds int) *Timer {
	return &Timer{
		ticker:          time.NewTicker(time.Second),
		done:            make(chan bool),
		durationSeconds: durationSeconds,
	}
}

// Start the timer and print the time left to the user
func (timer *Timer) StartTimer() {
	for {
		select {
		// If the timer's done channel emits a value (aka it is set to true)
		case <-timer.done:
			return
		// Every time the ticker channel emits a value (aka one second passes)
		case <-timer.ticker.C:
			timer.durationSeconds -= 1
			clearScreen()
			fmt.Println("Seconds left:", timer.durationSeconds)
		}
	}
}

// Stop the timer
func (timer *Timer) StopTimer() {
	timer.ticker.Stop()
	// Set done to true, which will tell the StartTimer() function to stop
	timer.done <- true
	clearScreen()
}
