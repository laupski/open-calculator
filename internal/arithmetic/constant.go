package arithmetic

import (
	"errors"
	"math"
)

type Constant struct {
	Value          string
	Representation interface{}
}

var constants = map[string]Constant{
	"pi": Constant{"pi", math.Pi},
	"e":  Constant{"e", math.E},
	"π":  {"π", math.Pi},
}

func NewConstant(c string) (*Constant, error) {
	if val := constants[c]; val.Value != "" {
		return &val, nil
	}
	return nil, errors.New("constant not found")
}

func (c Constant) String() string {
	return c.Value
}
