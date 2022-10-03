package goroutines

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func GenerateIDs(IDsChan chan<- string, wg *sync.WaitGroup) {

	for i := 0; i < 100; i++ {
		id := uuid.New()
		IDsChan <- fmt.Sprintf("%d. %s\n", i+1, id.String())
	}

	close(IDsChan)

	wg.Done()
}
func LogIds(IDsChan <-chan string, fakeIDsChan <-chan string, wg *sync.WaitGroup) {

	for id := range IDsChan {
		fmt.Print(id)
	}
	wg.Done()
}

func GenerateFakeIDs(fakeIDsChan chan<- string, wg *sync.WaitGroup) {

	for i := 0; i < 50; i++ {
		id := uuid.New()
		fakeIDsChan <- fmt.Sprintf("%d. %s\n", i+1, id.String())
	}

	close(fakeIDsChan)

	wg.Done()
}
