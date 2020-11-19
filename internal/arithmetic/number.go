package arithmetic

import (
	"errors"
	"fmt"
	"strconv"
)

type Number struct {
	Value          float64
	Representation string
}

func (n Number) String() string {
	return n.Representation
}

func NewNumber(i interface{}) (*Number, error) {
	switch i.(type) {
	case int:
		s := strconv.Itoa(i.(int))

		return &Number{
			Value:          float64(i.(int)),
			Representation: s,
		}, nil
	case string:
		if len(i.(string)) == 0 {
			return nil, errors.New("empty string input")
		}

		f, err := strconv.ParseFloat(i.(string), 64)
		if err == nil {
			return &Number{
				Value:          f,
				Representation: i.(string),
			}, nil
		}
	case float64:
		s := fmt.Sprintf("%f", i.(float64))

		return &Number{
			Value:          i.(float64),
			Representation: s,
		}, nil
	case *Constant:
		c := i.(*Constant)
		return &Number{
			Value:          c.Representation.(float64),
			Representation: c.Value,
		}, nil
	default:
		return nil, errors.New("broken input on Number")
	}

	return nil, errors.New("broken input on Number")
}
