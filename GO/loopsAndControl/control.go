package main

import (
	"errors"
	"fmt"
)

func EvenOrOdd(initNumber, finalNumber, iterate int) (*string, error) {
	if initNumber > 0 || finalNumber < 0 || iterate < 0 || finalNumber < initNumber {
		return nil, errors.New("The numbers can't be negative and the final number can't be less than the initial number")
	}

	count := initNumber
	text := ""
	for count <= finalNumber {
		if count%2 == 0 {
			text += fmt.Sprintf("%d is even\n", count)
		} else {
			text += fmt.Sprintf("%d is odd\n", count)
		}
		count += iterate
	}
	return &text, nil
}

func main() {
	result, err := EvenOrOdd(0, 50, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*result)
}
