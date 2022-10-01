package main

import "github.com/Erickype/GoChannels/goroutines"

func main() {
	IDsChan := make(chan string)
	goroutines.GenerateIDs(IDsChan)
}
