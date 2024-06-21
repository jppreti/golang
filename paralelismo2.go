package main

import (
	"fmt"
)

func processa(canal chan string) {
	canal <- "processado"
}

func main() {
	canal := make(chan string)
	go processa(canal)
	result := <-canal
	fmt.Println(result)
}
