package main

import "sync"

var racyAccess int
var wg sync.WaitGroup

// CriticalSection receives a copy of mutex .
func criticalSection(m sync.Mutex) {
	m.Lock()
	racyAccess++
	m.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	mutex := sync.Mutex{}
	// passes a copy of m to A .
	go criticalSection(mutex)
	go criticalSection(mutex)
	wg.Wait()
}
