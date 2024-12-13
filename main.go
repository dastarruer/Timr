package main

import (
	"fmt"
	"time"
    "os"
    "os/exec"
) 


type Timer struct {
    ticker *time.Ticker
    done chan bool
    duration int
}

func main() {
	// Create a ticker; this is a channel that is given a value every time.Second
    secondsLeft := 2
	timer := createTimer(secondsLeft)

    // Start the timer
    go func() {
        for {
            select {
            case <-timer.done:
                return
            // Every time the ticker channel emits a value (aka one second passes)
            case <-timer.ticker.C:
                secondsLeft -= 1
                clearScreen()  
                fmt.Println("Seconds left:", secondsLeft)
            }
        }
    }()

    time.Sleep(time.Duration(secondsLeft) * time.Second)

    // Stop the timer
    stopTimer(timer)
    fmt.Println("Time is up!")
}

func clearScreen() {
    cmd := exec.Command("clear") // For Linux, macOS
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func createTimer(duration int) *Timer {
    return &Timer {
        ticker: time.NewTicker(time.Second),
        done: make (chan bool),
        duration: duration,
    }
}

func stopTimer(timer *Timer) {
    timer.ticker.Stop()
    timer.done <- true
}
