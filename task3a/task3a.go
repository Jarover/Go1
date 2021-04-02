package main

import (
	"fmt"
	"log"
)

func main() {

	var number, digit int
	arr := [3]string{"Единиц", "Десятков", "Тысяч"}
	fmt.Println(" Введите число :")
	if _, err := fmt.Scanln(&number); err != nil {
		log.Fatalln(err)
	}
	if number < 1 || number > 999 {
		log.Fatalln("Некорректное число")
	}
	for i := 0; number > 0; i++ {
		digit = number % 10
		number = (number - digit) / 10
		fmt.Println(arr[i], " - ", digit)
	}
}
