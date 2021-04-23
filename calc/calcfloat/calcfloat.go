package calcfloat

// калькулятор реализующий работу в обратной польской нотации
// для чисел с плавающей запятой

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const helpMessage = "command format: 'num1 num2 cmd'"

var Calculate Calc

type Calc struct {
	a, b, res float64
	cmdMap    map[string]func() (interface{}, error)
}

func (c *Calc) Sqrt() (interface{}, error) {
	return math.Sqrt(c.b), nil
}

func (c *Calc) Sum() (interface{}, error) {
	return c.a + c.b, nil
}

func (c *Calc) Subtract() (interface{}, error) {
	return c.a - c.b, nil
}

func (c *Calc) Multiply() (interface{}, error) {

	return c.a * c.b, nil
}

func (c *Calc) Divide() (interface{}, error) {
	if c.b == 0 {
		return 0, fmt.Errorf("ERROR: dividing by zero")
	}
	return c.a / c.b, nil
}

func (c *Calc) ProcessCmd(cmd string) (interface{}, error) {
	//fmt.Println(c.a, c.b, cmd)
	if cmdFunc, ok := c.cmdMap[cmd]; ok {
		return cmdFunc()
	}
	return 0, fmt.Errorf("command not implemented")
}

func (c *Calc) InputString(input string) (interface{}, error) {
	var cmd string

	var err1, err2 error
	tokens := []string{}
	for _, token := range strings.Split(input, " ") {
		token = strings.TrimSpace(token) // отрезаем пробелы
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	//fmt.Println(c.res)

	switch len(tokens) {
	case 1:
		c.b = c.res
		cmd = tokens[0]
	case 2:
		cmd = tokens[1]
		c.a = c.res
		c.b, err2 = strconv.ParseFloat(tokens[0], 64)

	case 3:
		cmd = tokens[2]
		c.a, err1 = strconv.ParseFloat(tokens[0], 64)
		c.b, err2 = strconv.ParseFloat(tokens[1], 64)

	default:
		return 0, fmt.Errorf(helpMessage)

	}

	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("ERROR validate input")
	}
	res, err := c.ProcessCmd(cmd)
	if err != nil {
		return 0, err
	}
	c.res = res.(float64)

	return res, nil
}

func (c *Calc) DoCalc(input string) {
	fmt.Println(input)
	res, err := c.InputString(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("= ", res)
	}
}

// непрерывный ввод команд в калькулятор с клавиатуры
func (c *Calc) Work() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("Вводите операции вычисления в формате обратной польской нотации 'num1 num2 cmd' или 'num2 cmd' или 'cmd'")
	for scan.Scan() {
		inputText := scan.Text()
		if len(inputText) == 0 {
			return
		}

		res, err := c.InputString(inputText)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(" = ", res)
		}
	}
}

func (c *Calc) Init() {
	c.cmdMap = make(map[string]func() (interface{}, error))
	c.cmdMap["+"] = c.Sum
	c.cmdMap["-"] = c.Subtract
	c.cmdMap["*"] = c.Multiply
	c.cmdMap["/"] = c.Divide
	c.cmdMap["sqrt"] = c.Sqrt
}

func init() {
	Calculate.Init()

}
