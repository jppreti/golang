package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", tipo, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go contador("a")
	go contador("b")
	time.Sleep(time.Second * 20)
}
