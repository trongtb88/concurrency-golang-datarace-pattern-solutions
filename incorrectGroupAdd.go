package main

import (
	"fmt"
	"sync"
)

// ORIGINAL
// Using GroupAdd should outside and before goroutine
/*
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
*/

func main() {
	var wg sync.WaitGroup
	var racyAccess int
	wg.Add(1)
	go func(race *int) {
		*race++
		wg.Done()
	}(&racyAccess)
	wg.Wait()
	racyAccess++
	fmt.Println(racyAccess)
}
