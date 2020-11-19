package arithmetic

import (
	"errors"
	"fmt"
)

type Parenthesis struct {
	Value string
	Side  string
}

var parentheses = map[string]Parenthesis{
	"(": Parenthesis{"(", "left"},
	"[": Parenthesis{"(", "left"},
	"{": Parenthesis{"(", "left"},
	")": Parenthesis{")", "right"},
	"]": Parenthesis{"]", "right"},
	"}": Parenthesis{"}", "right"},
}

func (p Parenthesis) String() string {
	return fmt.Sprintf(p.Value)
}

func NewParenthesis(s string) (*Parenthesis, error) {
	if val := parentheses[s]; val.Value != "" {
		return &val, nil
	}

	return nil, errors.New("parenthesis not found")
}
