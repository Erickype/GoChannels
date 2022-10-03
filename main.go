package main

import (
	"github.com/Erickype/GoChannels/goroutines"
	"sync"
)

func main() {
	IDsChan := make(chan string)
	fakeIDsChan := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(3)

	go goroutines.GenerateIDs(IDsChan, wg)
	go goroutines.GenerateFakeIDs(fakeIDsChan, wg)
	go goroutines.LogIds(IDsChan, fakeIDsChan, wg)

	wg.Wait()
}
