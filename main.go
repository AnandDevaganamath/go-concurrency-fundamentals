package main

import (
	"sync"
)

var a string

// func f(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	print(a) 
// }

func main() {
	a = "hello"
	var wg sync.WaitGroup
	wg.Add(1)
	// go f(&wg)

	go func() {
		defer wg.Done()
		print(a)
	}()
	wg.Wait()
}
