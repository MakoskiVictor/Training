package main

import (
	"fmt"
	"time"
)

func main() {
	// Iniciar un contador de tiempo
	start := time.Now()

	for i := 0; i < 100; i++ {
		calculate := i * 3
		fmt.Println(calculate)
	}

	duration := time.Since(start)

	fmt.Println("Duration: ", duration)

	// Quitando el init and post statement (While loop)
	sum := 1
	for sum < 100 {
		sum += sum
	}

	fmt.Println("Sum: ", sum)
}
