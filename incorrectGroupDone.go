package main

import "sync"

var racyAccess int

func work(wg *sync.WaitGroup) int {
	defer wg.Done()
	return 42
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		racyAccess = work(&wg)
	}()
	wg.Wait()
	racyAccess++
}
