package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)


func main() {
	arr := []int{}
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("Вводите целые числа для наполнения массива")
	for scan.Scan() {
		inputText := scan.Text()
		if len(inputText) == 0 {
			break
		}
		if inputNumber, err := strconv.Atoi(inputText); err == nil {
			arr = append(arr, inputNumber)
			fmt.Println(arr)
		} else {
			fmt.Println(err)
		}
	}
	if err := scan.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	InsertionSort(arr)
}
// Функция сортировки вставками
func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		tempElement := ar[i]
		j := i
		for ; j > 0 && ar[j-1] > tempElement; j-- {
			ar[j] = ar[j-1]
			//fmt.Println(j, ar)
		}
		ar[j] = tempElement
		//fmt.Println(ar)
	}
	fmt.Println(ar)
}