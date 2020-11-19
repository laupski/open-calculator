package arithmetic

import "errors"

type Separator struct {
	Value string
}

var separatorMappings = map[string]Separator{
	",": Separator{","},
}

func NewSeparator(s string) (*Separator, error) {
	if val := separatorMappings[s]; val.Value != "" {
		return &val, nil
	}

	return nil, errors.New("separator not found")
}

func (s Separator) String() string {
	return s.Value
}
