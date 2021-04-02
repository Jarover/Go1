package main

import (
	"fmt"
	"log"
)

func main() {

	var number, hundreds, tens, ones int64

	fmt.Print("Введите трехзначное число: ")

	if _, err := fmt.Scanln(&number); err != nil {
		panic(err)

	}

	if number > 999 || number < 100 {
		log.Fatalln("Введенное число некорректно!")

	}

	ones = number % 10
	fmt.Printf("Eдиницы: %d\n", ones)

	tens = ((number % 100) - ones) / 10
	fmt.Printf("Десятки: %d\n", tens)

	hundreds = (number - tens*10 - ones) / 100
	fmt.Printf("Сотни: %d\n", hundreds)

}
