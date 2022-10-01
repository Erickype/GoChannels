package main

import (
	"github.com/Erickype/GoChannels/goroutines"
	"sync"
)

func main() {
	IDsChan := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go goroutines.GenerateIDs(IDsChan, wg)
	go goroutines.LogIds(IDsChan, wg)

	wg.Wait()
}
