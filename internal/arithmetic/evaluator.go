package arithmetic

import "errors"

func Evaluate(queue *TokenQueue) (*Token, error) {
	if queue.IsEmpty() {
		return nil, errors.New("empty queue")
	}

	numberStack := NewTokenStack()

	for !queue.IsEmpty() {
		token, _ := queue.Dequeue()
		if token.TokenType == operator {
			operation := token.Value.(*Operator)
			var numbers []*Number

			popped, err := numberStack.Pop()
			if err == nil && popped.TokenType == constant {
				num, _ := NewNumber(popped.Value.(*Constant))
				numbers = append(numbers, num)
			} else if err == nil && popped.TokenType == number {
				num := popped.Value.(*Number)
				numbers = append(numbers, num)
			}

			if operation.OperatorKey != negative && operation.OperatorKey != positive {
				popped, err = numberStack.Pop()
				if err == nil && popped.TokenType == constant {
					num, _ := NewNumber(popped.Value.(*Constant))
					numbers = append(numbers, num)
				} else if err == nil && popped.TokenType == number {
					numbers = append(numbers, popped.Value.(*Number))
				}
			}

			var answer *Number
			if len(numbers) == 1 {
				answer, err = operation.Function(numbers[0], nil)
				if err != nil {
					return nil, err
				}
			} else {
				answer, err = operation.Function(numbers[0], numbers[1])
				if err != nil {
					return nil, err
				}
			}

			newToken := NewToken(answer, number)
			numberStack.Push(newToken)
		} else if token.TokenType == function {
			function := token.Value.(*Function)
			var numbers []*Number

			if function.Arguments != 0 {
				for i := 1; i <= function.Arguments; i++ {
					popped, _ := numberStack.Pop()
					if popped.TokenType == constant {
						num, _ := NewNumber(popped.Value.(*Constant))
						numbers = append(numbers, num)
					} else {
						numbers = append(numbers, popped.Value.(*Number))
					}
				}
			} else {
				for !numberStack.IsEmpty() {
					popped, _ := numberStack.Pop()
					if popped.TokenType == constant {
						num, _ := NewNumber(popped.Value.(*Constant))
						numbers = append(numbers, num)
					} else {
						numbers = append(numbers, popped.Value.(*Number))
					}
				}
			}

			answer, err := function.Implementation(numbers...)
			if err != nil {
				return nil, err
			}
			newToken := NewToken(answer, number)
			numberStack.Push(newToken)
		} else if token.TokenType == number || token.TokenType == constant {
			numberStack.Push(token)
		}
	}

	answer, _ := numberStack.Pop()
	return answer, nil
}
