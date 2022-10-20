package main

import (
	"strconv"
	"sync"
)

func processOne(x string) string {
	return x
}

func processAll(uuids []string) int {
	var myResults []string
	var mutex sync.Mutex
	safeAppend := func(res string) {
		mutex.Lock()
		myResults = append(myResults, res)
		mutex.Unlock()
	}

	var wg sync.WaitGroup
	wg.Add(len(uuids))
	for _, uuid := range uuids {
		go func(id string, results []string) {
			res := processOne(id)
			safeAppend(res)
			wg.Done()
		}(uuid, myResults) // slice read without holding lock
	}
	wg.Wait()
	return 0
}

func main() {
	var uuids []string
	// Create 1000 uuids
	for i := 0; i < 1000; i++ {
		uuids = append(uuids, strconv.FormatInt(int64(i), 10))
	}
	_ = processAll(uuids[:])
}
