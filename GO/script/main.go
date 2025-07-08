package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var name string
	var age string

	fmt.Println("Ingresa tu nombre:")
	fmt.Scanln(&name)
	fmt.Println("Ingresa tu edad:")
	fmt.Scanln(&age)
	user_age, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Debías ingresar un número válido")
	} else {
		if user_age >= 18 {
			fmt.Print("¡Felicidades ", name, ", sos mayor de edad!")
		} else {
			fmt.Print("¡Ups ", name, ", parece que no sos mayor de edad todavía!")
		}
	}
	fmt.Println("\n", "La consola se cerrará en 5 segundos")
	time.Sleep(5 * time.Second)
}
