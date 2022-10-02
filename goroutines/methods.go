package goroutines

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func GenerateIDs(IDsChan chan string, wg *sync.WaitGroup) {

	for i := 0; i < 100; i++ {
		id := uuid.New()
		IDsChan <- fmt.Sprintf("%d. %d\n", i+1, id)
	}
	wg.Done()
}
func LogIds(IDsChan chan string, wg *sync.WaitGroup) {
	id := <-IDsChan
	fmt.Printf("Id: %v\n", id)
	wg.Done()
}
