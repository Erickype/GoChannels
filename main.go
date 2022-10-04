package main

import (
	"github.com/Erickype/GoChannels/goroutines"
	"sync"
)

func main() {
	IDsChan := make(chan string)
	fakeIDsChan := make(chan string)
	closedChan := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(3)

	go goroutines.GenerateIDs(IDsChan, closedChan, wg)
	go goroutines.GenerateFakeIDs(fakeIDsChan, closedChan, wg)
	go goroutines.LogIds(IDsChan, fakeIDsChan, wg, closedChan)

	wg.Wait()
}
