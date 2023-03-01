package parser

import (
	"fmt"

	"github.com/csirak1528/modlang/errors"

	"github.com/csirak1528/modlang/token"
)

type StackNode struct {
	val  any
	next *StackNode
}
type Stack struct {
	head *StackNode
	Size int
}

func (s *Stack) Push(t any) {
	cur := StackNode{val: t, next: s.head}
	s.Size++
	s.head = &cur
}

func (s *Stack) Pop() any {
	out := s.head
	if out != nil {
		s.head = out.next
		s.Size--
		return out.val
	}
	return nil
}

func (s *Stack) Peek() any {
	return s.head.next.val
}

func (s *Stack) Cur() any {
	return s.head.val
}

func (s *Stack) ToArray() any {
	arr := make([]any, s.Size)
	for s.Size > 0 {
		arr[s.Size] = s.Pop()
	}
	return arr
}

func (s *Stack) Parse() *Operation {
	parser := s.GetParse()
	return parser.Parse()
}

func (s *Stack) GetParse() *Parser {
	stackItems := s.ToArray().([]any)
	tokens := make([]*token.Token, len(stackItems))
	opStack := &Stack{}
	for i, item := range stackItems {
		switch item.(type) {
		case *token.Token:
			tokens[i] = item.(*token.Token)
		case *Operation:
			opStack.Push(item.(*Operation))
		default:
			panic(errors.StackSyntaxError)
		}
	}

	return &Parser{Tokens: tokens, ItemStack: opStack}
}
func (s *Stack) ParseAll() *Operation {
	parser := s.GetParse()
	stack := parser.ParseAll()
	return &Operation{Type: token.SCOPE, Children: []any{stack}}
}
func (s *Stack) Log(indent int) {
	arr := s.ToArray().([]any)
	for i, item := range arr {
		switch item.(type) {
		case *token.Token:
			fmt.Print(item.(*token.Token).GetLog() + " ")
		case *Operation:
			if item.(*Operation).Type == token.SCOPE && i != 0 {
				printIndent(indent)
			}
			item.(*Operation).Log(indent)
		}
	}
	for _, item := range arr {
		s.Push(item)
	}
}
