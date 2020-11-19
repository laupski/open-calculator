package arithmetic

import (
	"errors"
	"fmt"
	"github.com/laupski/open-calculator/internal/collection"
	"regexp"
	"strings"
)

const allowed = numbers + p + e + m + d + a + s + decimal + space + separator + alphabet + mod + mathConstants
const nonNumber = operators + separator + alphabet + p + mathConstants
const validNumber = numbers + decimal
const operators = e + m + d + a + s + mod
const numbers = "1234567890"
const decimal = "."
const p = leftParenthesis + rightParenthesis
const leftParenthesis = "([{"
const rightParenthesis = "}])"
const e = "^"
const m = "*xX×"
const d = "/÷"
const a = "+"
const s = "-−"
const mod = "%"
const space = " "
const separator = ","
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const mathConstants = "π"
const (
	operator          = iota // 0
	function                 // 1
	number                   // 2
	constant                 // 3
	leftparenthesis          // 4
	rightparenthesis         // 5
	argumentseparator        // 6
)

type TokenType int

type Token struct {
	Value     interface{}
	TokenType TokenType
}

type TokenQueue struct {
	contents []*Token
	length   int
}

type TokenStack struct {
	contents []*Token
	length   int
}

func NewToken(value interface{}, tokenType TokenType) *Token {
	return &Token{Value: value, TokenType: tokenType}
}

func (t *Token) String() string {
	return fmt.Sprintf("%v",t.Value)
}

func NewTokenQueue() *TokenQueue {
	return &TokenQueue{
		contents: make([]*Token, 0),
		length:   0,
	}
}

func (q *TokenQueue) Enqueue(t *Token) {
	q.contents = append(q.contents, t)
	q.length++
}

func (q *TokenQueue) Dequeue() (*Token, error) {
	if q.Len() < 1 {
		return nil, errors.New("empty queue")
	}

	dequeued := q.contents[0]
	q.contents = q.contents[1:]
	q.length--
	return dequeued, nil
}

func (q *TokenQueue) Len() int {
	return q.length
}

func (q *TokenQueue) Peek() (*Token, error) {
	if q.Len() < 1 {
		return nil, errors.New("empty queue")
	}

	return q.contents[0], nil
}

func (q *TokenQueue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *TokenQueue) String() string {
	var s string
	for _, c := range q.contents {
		switch c.TokenType {
		case number:
			s += c.Value.(*Number).String() + " "
		case operator:
			s += c.Value.(*Operator).String() + " "
		case leftparenthesis:
			s += c.Value.(*Parenthesis).String() + " "
		case rightparenthesis:
			s += c.Value.(*Parenthesis).String() + " "
		case constant:
			s += c.Value.(*Constant).String() + " "
		case function:
			s += c.Value.(*Function).String() + " "
		case argumentseparator:
			s += c.Value.(*Separator).String() + " "
		}
	}
	return s
}

func NewTokenStack() *TokenStack {
	return &TokenStack{
		contents: make([]*Token, 0),
		length:   0,
	}
}

func (s *TokenStack) Len() int {
	return s.length
}

func (s *TokenStack) Peek() (*Token, error) {
	if s.Len() < 1 {
		return nil, errors.New("empty stack")
	}
	return s.contents[s.length-1], nil
}

func (s *TokenStack) Pop() (*Token, error) {
	if s.Len() < 1 {
		return nil, errors.New("empty stack")
	}

	popped := s.contents[s.length-1]
	s.length--
	return popped, nil
}

func (s *TokenStack) Push(element *Token) {
	s.length++
	if len(s.contents) < s.length {
		// there is not enough space in s.inner. we need to append to the end of it to cause a re-allocation
		s.contents = append(s.contents, element)
	} else {
		// the s.inner slice already has enough space for this element
		s.contents[s.length-1] = element
	}
}

func (s *TokenStack) IsEmpty() bool {
	return s.Len() == 0
}

// unclean string input -> tokens for the algorithm to consume
func Tokenize(input string) ([]*Token, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid empty input")
	}

	if _, err := isClean(input); err != nil {
		return nil, errors.New("invalid characters found")
	}

	if !HasMatchingParentheses(input) {
		return nil, errors.New("invalid mismatched parameters")
	}

	tempInteger := ""
	tempAlphabet := ""
	var tokens []*Token

	// TODO add error index
	for i, r := range input {
		if strings.ContainsAny(string(r), space+separator+nonNumber) {
			debug := string(r)
			_ = debug
			if len(tempInteger) > 0 {
				newNumber, err := NewNumber(tempInteger)
				if err == nil {
					newToken := NewToken(newNumber, number)
					tokens = append(tokens, newToken)
					tempInteger = ""
					tempAlphabet = ""
				}
			}

			newFunction, err := NewFunction(tempAlphabet + string(r))
			if err == nil {
				newToken := NewToken(newFunction, function)
				tokens = append(tokens, newToken)
				tempInteger = ""
				tempAlphabet = ""
				continue
			}

			newConstant, err := NewConstant(tempAlphabet + string(r))
			if err == nil {
				newToken := NewToken(newConstant, constant)
				tokens = append(tokens, newToken)
				tempInteger = ""
				tempAlphabet = ""
				continue
			}

			newParenthesis, err := NewParenthesis(string(r))
			if err == nil {
				var newToken *Token
				if newParenthesis.Side == "left" {
					newToken = NewToken(newParenthesis, leftparenthesis)
				} else {
					newToken = NewToken(newParenthesis, rightparenthesis)
				}
				tokens = append(tokens, newToken)
				tempInteger = ""
				tempAlphabet = ""
				continue
			}

			newOperator, err := NewOperator(string(r))
			if err == nil {
				newToken := NewToken(newOperator, operator)
				tokens = append(tokens, newToken)
				tempInteger = ""
				tempAlphabet = ""
				continue
			}

			newSeparator, err := NewSeparator(string(r))
			if err == nil {
				newToken := NewToken(newSeparator, argumentseparator)
				tokens = append(tokens, newToken)
				tempInteger = ""
				tempAlphabet = ""
				continue
			}

			if strings.ContainsAny(string(r), alphabet) {
				tempAlphabet += string(r)
				continue
			}
		} else if strings.ContainsAny(string(r), validNumber) {
			if i == len(input)-1 {
				newNumber, err := NewNumber(tempInteger + string(r))
				if err == nil {
					newToken := NewToken(newNumber, number)
					tokens = append(tokens, newToken)
					tempInteger = ""
					tempAlphabet = ""
				}
			} else {
				tempInteger += string(r)
				continue
			}
		}
	}

	return tokens, nil
}

func isClean(s string) (bool, error) {
	cleaned := s
	for _, r := range allowed {
		cleaned = strings.ReplaceAll(cleaned, string(r), "")
	}
	if len(cleaned) > 0 {
		return false, errors.New("invalid input found: " + cleaned)
	}
	return true, nil
}

func HasMatchingParentheses(s string) bool {
	stack := collection.NewStack()
	reg := regexp.MustCompile("[^{}[\\]()]")
	input := reg.ReplaceAllString(s, "")
	openP := []rune{'(', '{', '['}
	closeP := []rune{')', '}', ']'}

	for _, r := range input {
		if IndexRune(openP, r) != -1 {
			stack.Push(r)
		} else {
			temp, e := stack.Pop()
			if e != nil {
				return false
			}

			if IndexRune(openP, temp) != IndexRune(closeP, r) {
				return false
			}
		}
	}

	return true
}

func IndexRune(list []rune, check interface{}) int {
	for i, v := range list {
		if v == check {
			return i
		}
	}

	return -1
}

func ToTokenQueue(tok []*Token) *TokenQueue {
	newQueue := NewTokenQueue()
	for _, t := range tok {
		newQueue.Enqueue(t)
	}
	return newQueue
}

// TODO
func (t *Token) compareTokens(u *Token) bool {
	if t.TokenType == u.TokenType {
		// TODO need to add comparisons between types
	}
	return false
}