package main

import (
	"fmt"
	"time"
)

//Thread 1
func main() {
	queue := make(chan int)

	//Thread 2
	go func() {
		i := 0
		for {
			queue <- i // Cheio
			i++
			time.Sleep(time.Second)
		}
	}()

	// Thread 1
	for x := range queue { // Esvaziar
		fmt.Println(x)
	}
}
