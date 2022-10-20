package main

import "sync"

type HealthGate struct {
	mutex sync.RWMutex
	ready bool
}

func cond() bool {
	return true
}
func (g *HealthGate) updateGate() {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	_ = g.ready
	if cond() {
		g.ready = true
	}
}

func main() {
	var h HealthGate
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		h.updateGate()
		wg.Done()
	}()
	go func() {
		h.updateGate()
		wg.Done()
	}()
	wg.Wait()
}
