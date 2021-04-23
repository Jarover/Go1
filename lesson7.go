package main

import (
	"fmt"

	"github.com/Jarover/Go1/calc"
	"github.com/Jarover/Go1/calc/calcfloat"
	"github.com/Jarover/Go1/calc/calcint"
	"github.com/Jarover/Go1/calc/calcuint"
)

func Calc(c calc.Calcer, input string) {
	c.DoCalc(input)
}

func main() {
	fmt.Println("Float test")
	f := new(calcfloat.Calc)
	f.Init()
	Calc(f, "2.4 3.3 +")
	Calc(f, "2.4 3.3 -")
	Calc(f, "-2.4 -3.3 -")
	Calc(f, "2.4 -3.3 -")
	Calc(f, "2.4 3.3 *")
	Calc(f, "2.4 3.3 /")
	Calc(f, "2.4 0 /")
	Calc(f, "2.4 sqrt")

	fmt.Println("Int test")
	i := new(calcint.Calc)
	i.Init()
	Calc(i, "2 3 +")
	Calc(i, "2 3 -")
	Calc(i, "2 3 -")
	Calc(i, "-2 -3 -")
	Calc(i, "2 -3 -")
	Calc(i, "2 3 *")
	Calc(i, "6 3 /")
	Calc(i, "2 0 /")
	Calc(i, "16 sqrt")

	fmt.Println("Uint test")
	u := new(calcuint.Calc)
	u.Init()
	Calc(u, "2 3 +")
	Calc(u, "2 3 -")
	Calc(u, "2 3 -")
	Calc(u, "-2 -3 -")
	Calc(u, "2 -3 -")
	Calc(u, "2 3 *")
	Calc(u, "6 3 /")
	Calc(u, "2 0 /")
	Calc(u, "16 sqrt")
}
