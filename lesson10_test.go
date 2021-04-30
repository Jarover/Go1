package main

import (
	"testing"

	"github.com/Jarover/Go1/calc/calcint"
	"github.com/stretchr/testify/assert"
)

func TestIntCalc(t *testing.T) {
	i := new(calcint.Calc)
	i.Init()

	tests := []struct {
		name     string
		op1, op2 string
		cmd      string
		want     float64
	}{
		{name: "1", op1: "4", op2: "2", cmd: "+", want: 6},
		{name: "2", op1: "4", op2: "2", cmd: "-", want: 2},
		{name: "3", op1: "4", op2: "3", cmd: "/", want: 0.75},
		{name: "4", op1: "4", op2: "2", cmd: "*", want: 8},
	}

	for _, tc := range tests {
		got, _ := i.InputString(tc.op1 + " " + tc.op2 + " " + tc.cmd)
		assert.Equal(t, tc.want, got, "they should be equal")

	}
}
