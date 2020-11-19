package arithmetic

import (
	"errors"
	"math"
)

const (
	addition       = iota // 0
	subtraction           // 1
	multiplication        // 2
	division              // 3
	exponential           // 4
	remainder             // 5
	negative              // 6
	positive              // 7
)

type OperatorKey int
type BinaryOp func(*Number, *Number) (*Number, error)

type Operator struct {
	Value         string
	Precedence    int
	Associativity string
	Function      BinaryOp
	OperatorKey
}

var operatorMappings = map[string]Operator{
	// substitutes for number flag
	"#": Operator{"-", 5, "right", negate, negative},
	"@": Operator{"+", 5, "right", reflectNum, positive},
	// normal operators
	"^": Operator{"^", 4, "right", power, exponential},
	"/": Operator{"/", 3, "left", divide, division},
	"%": Operator{"%", 3, "left", modulus, remainder},
	"*": Operator{"*", 3, "left", multiply, multiplication},
	"+": Operator{"+", 2, "left", add, addition},
	"-": Operator{"-", 2, "left", subtract, subtraction},
	//alternatives ones to include
	"x": Operator{"x", 3, "left", multiply, multiplication},
	"X": Operator{"X", 3, "left", multiply, multiplication},
	"×": Operator{"×", 3, "left", multiply, multiplication},
	"÷": Operator{"÷", 3, "left", divide, division},
	"−": Operator{"−", 2, "left", subtract, subtraction},
}

func NewOperator(s string) (*Operator, error) {
	if val := operatorMappings[s]; val.Value != "" {
		return &val, nil
	}

	return nil, errors.New("operator not found")
}

func negate(a *Number, b *Number) (*Number, error) {
	return NewNumber(a.Value * -1)
}

func reflectNum(a *Number, b *Number) (*Number, error) {
	return NewNumber(a.Value)
}

func add(a *Number, b *Number) (*Number, error) {
	if b == nil {
		return NewNumber(a.Value)
	}
	return NewNumber(a.Value + b.Value)
}

func subtract(a *Number, b *Number) (*Number, error) {
	if b == nil {
		return NewNumber(a.Value * -1)
	}
	return NewNumber(b.Value - a.Value)
}

func multiply(a *Number, b *Number) (*Number, error) {
	return NewNumber(a.Value * b.Value)
}

func divide(a *Number, b *Number) (*Number, error) {
	return NewNumber(b.Value / a.Value)
}

func modulus(a *Number, b *Number) (*Number, error) {
	return NewNumber(math.Mod(a.Value, b.Value))
}

func power(a *Number, b *Number) (*Number, error) {
	return NewNumber(math.Pow(b.Value, a.Value))
}

func (o *Operator) comparePrecedence(c Operator) bool {
	return o.Precedence >= c.Precedence
}

func (o Operator) String() string {
	return o.Value
}
