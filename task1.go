package main

import (
	"fmt"
)

func main() {

	var width, height, square int

	fmt.Print("Введите ширину прямоугольника: ")

	if _, err := fmt.Scanln(&width); err != nil {
		panic(err)

	}
	fmt.Print("Введите высоту прямоугольника: ")
	if _, err := fmt.Scanln(&height); err != nil {
		panic(err)
	}
	square = width * height
	fmt.Printf("Площадь прямоугольника: %d\n", square)
}
