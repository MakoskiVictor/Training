package main

import (
	"fmt"
	"time"
)

func main() {
	// Iniciar un contador de tiempo
	start := time.Now()

	for i := 0; i < 100000; i++ {
		calculate := i * 3
		fmt.Println(calculate)
	}

	duration := time.Since(start)

	fmt.Println("Duration: ", duration)
}
