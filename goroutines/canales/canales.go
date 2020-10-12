package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go goroutineChannel(i, c)
	}

	for i := 0; i < 5; i++ {
		rutina := <-c
		fmt.Printf("La rutina, numero %d, finalizo su tarea!\n", rutina)
	}
	close(c)
}

func goroutineChannel(numero int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("La rutina, numero %d, se esta ejecutando\n", numero)
	c <- numero
}

