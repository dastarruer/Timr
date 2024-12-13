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

func createTimer(durationSeconds int) *Timer {
	return &Timer{
		ticker:          time.NewTicker(time.Second),
		done:            make(chan bool),
		durationSeconds: durationSeconds,
	}
}

// Only works for linux/mac
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func startTimer(timer *Timer, durationSeconds *int) {
	for {
		select {
		case <-timer.done:
			return
		// Every time the ticker channel emits a value (aka one second passes)
		case <-timer.ticker.C:
			*durationSeconds -= 1
			clearScreen()
			fmt.Println("Seconds left:", *durationSeconds)
		}
	}
}

func stopTimer(timer *Timer) {
	timer.ticker.Stop()
	timer.done <- true
}