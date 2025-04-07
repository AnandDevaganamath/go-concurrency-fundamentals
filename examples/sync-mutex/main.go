package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex
    var x int

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            mu.Lock()
            x++ // Now safe
            mu.Unlock()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Final x:", x) // Always 1000
}