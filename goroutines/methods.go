package goroutines

import "github.com/google/uuid"

func GenerateIDs(IDsChan chan string) {
	id := uuid.New()
	IDsChan <- id.String()
}
