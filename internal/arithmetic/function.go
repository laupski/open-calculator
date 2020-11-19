package arithmetic

import (
	"errors"
	"math"
)

type VariadicOp func(...*Number) (*Number, error)

type Function struct {
	Value          string
	Implementation VariadicOp
	Arguments      int // if 0, load all tokens
}

var Functions = map[string]Function{
	"sin":  Function{"sin", sin, 1},
	"cos":  Function{"cos", cos, 1},
	"tan":  Function{"tan", tan, 1},
	"abs":  Function{"abs", abs, 1},
	"max":  Function{"max", max, 0},
	"min":  Function{"min", min, 0},
	"neg":  Function{"neg", neg, 1},
	"avg":  Function{"avg", avg, 2},
	"vavg": Function{"vavg", avg, 0},
	"sqrt": Function{"sqrt", sqrt, 1},
}

func NewFunction(f string) (*Function, error) {
	if val := Functions[f]; val.Value != "" {
		return &val, nil
	}

	return nil, errors.New("function not found")
}

func (f Function) String() string {
	return f.Value
}

func sin(n ...*Number) (*Number, error) {
	return NewNumber(math.Sin(n[0].Value))
}

func cos(n ...*Number) (*Number, error) {
	return NewNumber(math.Cos(n[0].Value))
}

func tan(n ...*Number) (*Number, error) {
	return NewNumber(math.Tan(n[0].Value))
}

func abs(n ...*Number) (*Number, error) {
	return NewNumber(math.Abs(n[0].Value))
}

func max(n ...*Number) (*Number, error) {
	max := n[0]
	for _, arg := range n {
		if arg.Value > max.Value {
			max = arg
		}
	}

	return max, nil
}

func min(n ...*Number) (*Number, error) {
	min := n[0]
	for _, arg := range n {
		if arg.Value < min.Value {
			min = arg
		}
	}

	return min, nil
}

func neg(n ...*Number) (*Number, error) {
	return NewNumber(- n[0].Value)
}

func avg(n ...*Number) (*Number, error) {
	var sum float64
	for _, num := range n {
		sum += num.Value
	}
	return NewNumber(sum / float64(len(n)))
}

func sqrt(n ...*Number) (*Number, error) {
	return NewNumber(math.Sqrt(n[0].Value))
}
