package main

import (
	"errors"
	"fmt"
)

// Con control de errores
func add(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, errors.New("Negative numbers are not allowed")
	}
	return x + y, nil
}

// Con variables nombradas
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	result, err := add(50, 25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(split(17))
}
