package main

import (
	"github.com/Jarover/Go1/calc"
	"github.com/Jarover/Go1/calc/calcfloat"
)

func Calc(c calc.Calcer, input string) {
	c.DoCalc(input)
}

func main() {

	f := new(calcfloat.Calc)
	f.Init()
	Calc(f, "2.4 3.3 +")

}
