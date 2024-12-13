package main

import (
	"fmt"
	"time"
) 

func main() {
	// Create a ticker; this is a channel that is given a value every 'd' 
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)



    go func() {
        for {
            select {
            case <-done:
                return
            // Every time the ticker channel emits a value (aka one second passes)
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(10 * time.Second)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")

}