package main

import (
	"fmt"
	"time"
    "os"
    "os/exec"
) 

func main() {
	// Create a ticker; this is a channel that is given a value every 'd' 
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
    secondsLeft := 10

    go func() {
        for {
            select {
            case <-done:
                return
            // Every time the ticker channel emits a value (aka one second passes)
            case <-ticker.C:
                secondsLeft -= 1
                cmd := exec.Command("clear") // For Linux, macOS
                cmd.Stdout = os.Stdout
                cmd.Run()
                fmt.Println("Seconds left:", secondsLeft)
            }
        }
    }()

    time.Sleep(10 * time.Second)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")

}