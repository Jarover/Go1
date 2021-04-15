package main

import (
	c "Go1/calc"
	"Go1/fibonachy"
	"fmt"
	"log"
	"time"
)

func main()  {

	var number int64
	fmt.Println("Тестируем алгоритмы вычисления чисел фибоначи")
	fmt.Print("Введите число ")

	if _, err := fmt.Scanln(&number); err != nil {
		panic(err)
	}

	start := time.Now()
	fmt.Println(fibonachy.Fibre(number))
	duration("Рекурсивный алгоритм",start)

	start = time.Now()
	fmt.Println(fibonachy.Fibmap(number))
	duration("Алгоритм с использованием map",start)


	// тестируем типовые операции калькулятора
	fmt.Println("тестируем типовые операции калькулятора")
	c.Calculate.DoCalc("11 2.4 +")
	c.Calculate.DoCalc("11 2.4 -")
	c.Calculate.DoCalc("11 2 *")
	c.Calculate.DoCalc("11 2 /")
	c.Calculate.DoCalc("11 0 /")
	c.Calculate.DoCalc("16 sqrt")
	c.Calculate.DoCalc("sqrt")

	// вводим значения с клавиатуры
	c.Calculate.Work()

}


// Вычисляем продолжительность со старта события
func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}