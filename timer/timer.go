package timer

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
	"timr/tui"
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

func convertToDigitalFormat(timeSeconds int) string {
	minutes := strconv.Itoa(timeSeconds / 60)
	seconds := strconv.Itoa(timeSeconds % 60)
	result := fmt.Sprintf("%s:%s", minutes, seconds)
	return result
}

// Return a pointer to a new Timer instance
func CreateTimer(durationSeconds int) *Timer {
	return &Timer{
		ticker:          time.NewTicker(time.Second),
		done:            make(chan bool),
		durationSeconds: durationSeconds,
	}
}

// Clear the screen and print the time left on the timer
func (timer *Timer) printTimeLeft() {
	clearScreen()
	tui.ShowTimer(convertToDigitalFormat(timer.durationSeconds))
}

// Start the timer and print the time left to the user
func (timer *Timer) StartTimer() {
	// We print the time left at the beginning, so that the timer does not start one second after the user starts the program
	timer.printTimeLeft()
	for {
		select {
		// If the timer's done channel emits a value (aka it is set to true)
		case <-timer.done:
			return
		// Every time the ticker channel emits a value (aka one second passes)
		case <-timer.ticker.C:
			timer.durationSeconds -= 1
			timer.printTimeLeft()
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
