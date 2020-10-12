package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go goroutine(i)
	}
	time.Sleep(4 * time.Second)
}

func goroutine(numero int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("La Goroutine, numero %d, termino su tarea\n", numero)
}

