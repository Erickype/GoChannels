package goroutines

import (
	"fmt"
	"github.com/google/uuid"
)

func GenerateIDs(IDsChan chan string) {
	id := uuid.New()
	IDsChan <- id.String()
}
func LogIds(IDsChan chan string) {
	id := <-IDsChan
	fmt.Printf("Id: %v\n", id)
}
