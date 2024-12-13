package main

import (
	"fmt"
	"time"
    "os"
    "os/exec"
) 

func main() {
	// Create a ticker; this is a channel that is given a value every time.Second
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
    secondsLeft := 10

    // Start the timer
    go func() {
        for {
            select {
            case <-done:
                return
            // Every time the ticker channel emits a value (aka one second passes)
            case <-ticker.C:
                secondsLeft -= 1
                clearScreen()  
                fmt.Println("Seconds left:", secondsLeft)
            }
        }
    }()

    time.Sleep(time.Duration(secondsLeft) * time.Second)

    // Stop the timer
    ticker.Stop()
    done <- true
    fmt.Println("Time is up!")

}

func clearScreen() {
    cmd := exec.Command("clear") // For Linux, macOS
    cmd.Stdout = os.Stdout
    cmd.Run()
}
