package main

import (
    "fmt"
    "time"
)

func printNumbers(id int) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Goroutine %d: %d\n", id, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    for i := 0; i < 3; i++ {
        go printNumbers(i)
    }
    time.Sleep(1 * time.Second) // Wait for goroutines
}