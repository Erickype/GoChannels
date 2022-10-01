package goroutines

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func GenerateIDs(IDsChan chan string, wg *sync.WaitGroup) {
	id := uuid.New()
	IDsChan <- id.String()
	wg.Done()
}
func LogIds(IDsChan chan string, wg *sync.WaitGroup) {
	id := <-IDsChan
	fmt.Printf("Id: %v\n", id)
	wg.Done()
}
