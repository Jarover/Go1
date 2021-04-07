package task2

import (
	"fmt"
)



// генератор простых чисел
func SimpleNumbers(number int) []int {


	var simple []int
	var flag bool


	for i := 2; i < number; i++ {
		flag = true
		for j:= 0;j< len(simple);j++ {
			if i%simple[j] == 0 {
				flag = false
				break
			}
		}
		if flag {
			simple = append(simple,i)
		}
	}



	return simple
}

func InputNumber()  {

	var number int
	fmt.Print("Введите число ")

	if _, err := fmt.Scanln(&number); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", SimpleNumbers(number))
}
