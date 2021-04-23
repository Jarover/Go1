package calc

type Calcer interface {
	DoCalc(string)
	Sum() (interface{}, error)
	Sqrt() (interface{}, error)
	Subtract() (interface{}, error)
	Multiply() (interface{}, error)
	Divide() (interface{}, error)
}
