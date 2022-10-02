package goroutines

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func GenerateIDs(IDsChan chan string, wg *sync.WaitGroup) {

	for i := 0; i < 100; i++ {
		id := uuid.New()
		IDsChan <- fmt.Sprintf("%d. %s\n", i+1, id.String())
	}
	wg.Done()
}
func LogIds(IDsChan chan string, wg *sync.WaitGroup) {

	for id := range IDsChan {
		fmt.Print(id)
	}
	wg.Done()
}
