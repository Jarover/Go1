package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		square, diameter, circle float64
	)

	fmt.Print("Введите плащадь круга: ")

	if _, err := fmt.Scanln(&square); err != nil {
		panic(err)

	}

	diameter = math.Sqrt(square/math.Pi) * 2

	fmt.Printf("Диаметр: %f\n", diameter)

	circle = diameter * math.Pi

	fmt.Printf("Окружность: %f\n", circle)
}
