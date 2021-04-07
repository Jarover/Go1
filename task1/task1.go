package task1

import (
	"fmt"
	"log"
	"os"
)

// Простой калькулятор
func Calc() {

	var number1, number2, res float32
	var op string
	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scanln(&number1); err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scanln(&number2); err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	if _, err := fmt.Scanln(&op); err != nil {
		log.Fatalln(err)
	}
	switch op {
	case "+":
		res = number1 + number2
	case "-":
		res = number1 - number2
	case "*":
		res = number1 * number2
	case "/":
		if number2 == 0 {
			fmt.Println("Нельзя делить на ноль")
			os.Exit(1)
		}
		res = number1 / number2
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %g\n", res)
}
