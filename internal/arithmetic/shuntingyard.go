package arithmetic

type ShuntingYard struct {
	infixTokens *TokenQueue
}

func NewShuntingYard(infixTokens *TokenQueue) *ShuntingYard {
	return &ShuntingYard{infixTokens: infixTokens}
}

func (sy *ShuntingYard) InfixToPostFix() *TokenQueue {
	outputQueue := NewTokenQueue()
	operatorStack := NewTokenStack()
	var lastTokenType TokenType = -1

	for !sy.infixTokens.IsEmpty() {
		currentToken, _ := sy.infixTokens.Dequeue()

		if currentToken.TokenType == number || currentToken.TokenType == constant {
			outputQueue.Enqueue(currentToken)
		} else if currentToken.TokenType == argumentseparator {
			peek, err := operatorStack.Peek()
			for err == nil && peek.TokenType != leftparenthesis {
				popped, _ := operatorStack.Pop()
				outputQueue.Enqueue(popped)
				peek, err = operatorStack.Peek()
			}
		} else if currentToken.TokenType == operator || currentToken.TokenType == function {
			peek, err := operatorStack.Peek()
			if err != nil {
				lastTokenType = currentToken.TokenType
				operatorStack.Push(currentToken)
				continue
			}

			// Take into consideration the ++ and +- and *- and *+ tokens etc.
			if currentToken.TokenType != function && (operatorStack.Len() == 0 && outputQueue.Len() == 0 ||
				lastTokenType == operator || lastTokenType == leftparenthesis) {
				if currentToken.Value.(*Operator).OperatorKey == subtraction {
					newOperator, _ := NewOperator("#") // unary minus operator
					currentToken.Value = newOperator
				} else if currentToken.Value.(*Operator).OperatorKey == addition {
					newOperator, _ := NewOperator("@") // unary reflection operator
					currentToken.Value = newOperator
				}
			}

			op, _ := currentToken.Value.(*Operator)

			for (op != nil && peek.TokenType == operator && (peek.Value.(*Operator).Precedence > op.Precedence) ||
				(op != nil && peek.TokenType == operator && peek.Value.(*Operator).Precedence == op.Precedence && op.Associativity == "left") ||
				(peek.TokenType == function)) && !operatorStack.IsEmpty() {
				popped, _ := operatorStack.Pop()
				outputQueue.Enqueue(popped)
				peek, err = operatorStack.Peek()
				if err != nil {
					break
				}
			}

			operatorStack.Push(currentToken)
		} else if currentToken.TokenType == leftparenthesis {
			operatorStack.Push(currentToken)
		} else if currentToken.TokenType == rightparenthesis {
			peek, _ := operatorStack.Peek()

			for peek.TokenType != leftparenthesis {
				popped, _ := operatorStack.Pop()
				outputQueue.Enqueue(popped)
				peek, _ = operatorStack.Peek()
			}

			if peek.TokenType == leftparenthesis {
				operatorStack.Pop()
			}

		}

		lastTokenType = currentToken.TokenType
	}

	for !operatorStack.IsEmpty() {
		popped, _ := operatorStack.Pop()
		outputQueue.Enqueue(popped)
	}

	return outputQueue
}
