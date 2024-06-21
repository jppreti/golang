package main

import (
	"fmt"
	"time"
)

func workerCalculaAlgoDificil(workerID string, msg chan int) {
	for res := range msg {
		fmt.Printf("Worker %s recebeu %d\n", workerID, res)
		time.Sleep(time.Second)
	}
}

//Thread 1
func main() {
	msg := make(chan int)

	go workerCalculaAlgoDificil("Joao", msg)
	go workerCalculaAlgoDificil("Maria", msg)

	for i := 0; i < 100; i++ {
		msg <- i
	}
}
