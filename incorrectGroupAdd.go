package main

import "sync"

func main() {
	var wg sync.WaitGroup
	var racyAccess int
	go func() {
		wg.Add(1)
		racyAccess++
		wg.Done()
	}()
	wg.Wait()
	racyAccess++
}
