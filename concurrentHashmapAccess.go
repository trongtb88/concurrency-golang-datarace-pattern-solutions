package main

import (
	"errors"
	"sync"
)

func processOne(x string) (string, error) {
	return x, errors.New("random error")
}

func combineErrors(errMap map[string]error) int {
	return 10
}

func processOrders(uuids []string) int {
	var errMap = make(map[string]error)
	var wg sync.WaitGroup
	wg.Add(len(uuids))
	for _, uuid := range uuids {
		go func(uuid string) {
			defer wg.Done()
			_, err := processOne(uuid)
			if err != nil {
				errMap[uuid] = err
				return
			}
		}(uuid)
	}
	wg.Wait()
	return combineErrors(errMap)
}

func main() {
	uuids := [2]string{"abcd", "efgh"}
	_ = processOrders(uuids[:])
}
